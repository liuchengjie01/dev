package login

import "github.com/gin-gonic/gin"

func Login(c *gin.Context)  {
	c.JSON(200, gin.H{
		"login" : " login success",
	})
}