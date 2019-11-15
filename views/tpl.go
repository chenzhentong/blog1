package views

import (
	"beeblog/models"
	"time"
)


func Add(a int,b int) int{
	return a+b
}
func IsArrayNull(comments []models.Comment) bool{
	if len(comments)==0{
		return true
	}
	return false
}
func IsStringNull(str string) bool{
	if str==""{
		return true
	}
	return false
}
func TimeParseString(date time.Time) string{
	return date.Format("2006-01-02 15:04:05")
}
func JudgeTopicType(sid string) int{
	if sid=="cid"{
		return 1
	}else if sid=="mid"{
		return 2
	}
	return 0
}
func JudgeCategory(str string) bool{
	if str=="category"{
		return true
	}
	return false
}
func JudgeMonth(str string) bool{
	if str=="month"{
		return true
	}
	return false
}