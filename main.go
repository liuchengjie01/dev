package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/services/creater"
	"server/services/login"
	"server/services/participate"
)
func main()  {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Successfully Run   --own by liuchengjie")
	})
	router.GET("/login", login.Login)

	stu := router.Group("/student")
	stu.POST("/participate/:test_node", participate.Participate)

	tea := router.Group("/teacher")
	tea.POST("/creater", creater.Create)
	tea.DELETE("/delete", creater.Delete)
	err := router.Run(":8729")
	if err != nil {
		logrus.WithError(err).Error(err)
	}
}
