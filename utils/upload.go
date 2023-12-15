package utils

import (
	"bytes"
	"errors"
	"fmt"
	"kel15/storage"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func GetFileUpload(c *gin.Context, isImageRequired bool) (multipart.File, *multipart.FileHeader, *bool, error) {
	maxSize := int64(1024000) // allow only 1MB of file size

	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		return nil, nil, nil, errors.New("Image too large")
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return nil, nil, &isImageRequired, errors.New("Could not get uploaded file")
	}

	return file, fileHeader, &isImageRequired, nil
}

func UploadToS3(user_id int, s *storage.StorageS3Stuct, file multipart.File, fileHeader *multipart.FileHeader, tempFileName string) (string, error) {
	defer file.Close()

	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s.StorageS3).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s.BucketName),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(fileHeader.Header.Get("Content-Type")),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	regilion := os.Getenv("AWS_REGION")
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.BucketName, regilion, tempFileName)

	return url, err
}
