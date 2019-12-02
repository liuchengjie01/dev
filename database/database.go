package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

//network data transmit
type Exam struct {
	Id int64 `gorm:"size:10;primary_key;AUTO_INCREMENT;not null" json:"id"`//考试ID
	FounderId uint64 `json:"founder_id"`//创建者Id
	TestTime []time.Time `json:"test_time"`//考试时间
	TestCode []string `gorm:"not null" json:"test_code"`//考试进入码
	TestName string `json:"test_name"`	//考试科目
	TestContents uint32 `json:"test_contents"`	//考试内容
	TestLocation []string `json:"test_location"`	//考试地点
	TestTeacher []string `json:"test_teacher"`	//复试老师
	TestAllPeople uint32 `json:"test_all_people"` //考试容纳人数
}

//maybe db data
type Test struct {
	FounderId uint64//考试ID
	TestTime time.Time//考试时间
	TestCode string//考试进入码
	TestName string//考试科目
	TestContents uint32//考试内容
	TestLocation string//考试地点
	TestTeacher []string//复试老师
}
type User struct {
	ID       uint64 `gorm:"column:ID;primary_key;AUTO_INCREMENT;not null"`
	Username string `gorm:"column:username;not null"`
	Pwd      string `gorm:"column:pwd;not null"`
}

var (
	Passport *gorm.DB
)
//初始化MySQL
func InitMySQL() error {
	var err error
	Passport, err = gorm.Open("mysql", "root:root1@tcp(ip:3306)/database?charset=utf8mb4&parseTime=True&loc=Local&readTimeout=500ms")
	if err != nil {
		return err
	}
	return nil
}

//通过考试码查找考试信息
func QueryExam(TestCode int64)(*Exam, error){
	var exam *Exam
	d := Passport.Where("test_code = ?", TestCode).First(exam)
	if d.Error != nil {
		return exam, d.Error
	}
	return exam, nil
}

//插入一场考试
func InsertExam(exam *Exam) error{
	d := Passport.Create(exam)
	if d.Error != nil {
		return d.Error
	}
	return nil
}

//删除一场考试
func DeleteExam(Id int64)  (error, bool){
	var exam Exam
	d := Passport.Where("id = ?", Id).First(&exam)
	if d.Error != nil {
		return d.Error , false
	}
	d = Passport.Delete(exam)
	if d.Error != nil {
		return d.Error, false
	}
	return nil, true
}


