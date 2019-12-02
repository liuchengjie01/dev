package participate

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/database"
	"strconv"
)

func Participate(c *gin.Context)  {
	log := logrus.New()
	testcode := c.Param("test_code")
	testcodeInt64, err := strconv.ParseInt(testcode, 10, 64)
	if err != nil {
		log.Errorf("error is %v", err)
		c.JSON(http.StatusInternalServerError, "Server error happened!")
	}
	test, err := database.QueryExam(testcodeInt64)
	if err != nil {
		log.Errorf("error is %v", err)
		c.JSON(http.StatusNotFound, "请求失败！")
		return
	}
	c.HTML(http.StatusOK, "test", test)
}

