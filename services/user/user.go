package user

import (
	"github.com/gin-gonic/gin"
	"server/database"
)

func Register(c *gin.Context)  {
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	err := database.AddUser(&database.User{
		Username: username,
		Pwd:      pwd,
	})
	if err != nil {
		c.JSON(404, "发生了未知错误")
		return
	}
}
func Login(c *gin.Context)  {
	inputUserName := c.PostForm("username")
	inputPwd := c.PostForm("password")
	getPwd, err := database.QueryPwdByName(inputUserName)
	if err != nil {
		c.JSON(404, "发生了未知错误")
		return
	}
	if getPwd=="RecordNotFound" {
		c.JSON(200, "账号密码错误")
	}
	if getPwd==inputPwd {
		c.JSON(200, gin.H{
			"user" : " user success",
		})
		return
	}
}