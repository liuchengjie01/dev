package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/database"
	"server/services/creater"
	"server/services/participate"
	"server/services/user"
)
func main()  {
	var err error
	err = nil
	err = database.InitMySQL()
	if err != nil {
		logrus.WithError(err).Errorf("err is %v", err)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Successfully Run   --own by liuchengjie")
	})
	passport := router.Group("/user")
	passport.POST("/login", user.Login)
	passport.POST("/register", user.Register)

	stu := router.Group("/student")
	stu.POST("/participate/:test_node", participate.Participate)

	tea := router.Group("/teacher")
	tea.POST("/creater", creater.Create)
	tea.DELETE("/delete", creater.Delete)
	err := router.Run(":8729")
	if err != nil {
		logrus.WithError(err).Errorf("err is %v", err)
	}
}
