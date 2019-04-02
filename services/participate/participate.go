package participate

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/database"
)

func SearchTest(c *gin.Context)  {
	log := logrus.New()
	testcode := c.Param("test_code")
	test, err := database.QueryExam(testcode)
	if err != nil {
		log.Errorf("error is %v", err)
		c.JSON(http.StatusNotFound, "请求失败！")
		return
	}
	c.HTML(http.StatusOK, "test", test)
}

func Participate(c *gin.Context)  {
	
}
