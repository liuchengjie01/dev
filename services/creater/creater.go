package creater

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"server/database"
	"strconv"
	"time"
)

func Create(c *gin.Context)  {
	log := logrus.New()
	m := make(map[string]interface{})
	req, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(req, &m)
	if err != nil {
		log.Errorf("error is %v", err)
		c.JSON(http.StatusBadRequest, "请求错误！")
		return
	}
	if len(m) == 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"message" : "request is not json"})
	}
	//以下取字段需要进行判空处理，目前暂时先不管
	id, err := strconv.ParseInt(m["id"].(string), 10, 64)
	if err != nil {
		log.Errorf("error is %v", err)
		return
	}
	tmp_founder_id, err:= strconv.ParseInt(m["founderid"].(string), 10, 64)
	if err != nil{
		log.Errorf("error is %v", err)
		c.JSON(http.StatusBadGateway,gin.H{
			"message" : "error",
		})
		return
	}
	founderid := uint64(tmp_founder_id)
	//testcode,err := redis.Strings(m["testcode"], nil)
	if err != nil {
		log.Errorf("error is %v", err)
		c.JSON(http.StatusBadGateway, gin.H{
			"message" : "toSlice failed",
		})
		return
	}
	testtime := m["testtime"].([]time.Time)
	testname := m["testname"].(string)
	tmpContents, err := strconv.ParseInt(m["testcontents"].(string), 10, 32)
	if err != nil{
		log.Errorf("error is %v", err)
		return
	}
	testcontents := uint32(tmpContents)
	//testlocation, err := redis.Strings(m["testlocation"],nil)
	if err != nil {
		log.Errorf("error is %v", err)
		c.JSON(http.StatusBadGateway, gin.H{
			"message" : "toSlice failed",
		})
		return
	}
	var testcode []string
	testcode = m["testcode"].([]string)
	testlocation := m["testlocation"].([]string)
	exam := database.Exam{
		Id           : id,
		FounderId    : founderid,
		TestCode     : testcode,
		TestTime     : testtime,	//多个时间
		TestName     : testname,
		TestContents : testcontents,
		TestLocation : testlocation,
	}
	err = database.InsertExam(&exam)
	if err != nil{
		log.Errorf("error is %v", err)
		c.JSON(http.StatusBadGateway,gin.H{
			"message" : "insert failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"Message" : "OK",
	})
}

func Delete(c *gin.Context)  {
	log := logrus.New()
	//此处要进行权限验证，并非所有人都能随便删除考试,目前开发人员不够，先不写
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Errorf("error is %v", err)
		return
	}
	err, ok := database.DeleteExam(id)
	if err != nil || ok != true {
		fmt.Println("delete failed")
		c.JSON(403, gin.H{
			"Delete" : "Failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"Message" : "OK",
	})
}
