package handler

import "github.com/gin-gonic/gin"

func (handler *Handler) CheckToken(c *gin.Context) {
	_, err := handler.GetUserByToken(c)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
			"code":    401,
			"success": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "",
		"code":    200,
		"success": true,
	})
	return

}
