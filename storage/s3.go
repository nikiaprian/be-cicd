package storage

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type StorageS3Stuct struct {
	StorageS3  *session.Session
	BucketName string
}

func NewS3() *StorageS3Stuct {
	key := os.Getenv("AWS_ACCESS_KEY_ID")
	secret := os.Getenv("AWS_SECRET_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			key,    // id
			secret, // secret
			""),    // token can be left blank for now
	})
	if err != nil {
		return nil
	}

	return &StorageS3Stuct{
		StorageS3:  s,
		BucketName: bucket,
	}
}
