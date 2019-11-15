package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db orm.Ormer



type Category struct{
	Id int64
	Title string
	Created time.Time `orm:"index"`
	Views int64 `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
	Topics []*Topic `orm:"rel(m2m)"`


}
type Topic struct{
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views int64 `orm:"index"`
	Author string
	ReplyTime time.Time `orm:"index"`
	ReplyCount int64
	ReplyLastUserId int64
	Category *Category `orm:"rel(fk)"`
	Comments []*Comment `orm:"reverse(many)"`
}
type Comment struct{
	Id int64

	Name string
	Email string
	Website string
	Content string 	`orm:"size(1000)"`
	CreatedTime time.Time `orm:"index"`
	Topic *Topic  `orm:"rel(fk)"`
}
type TopicRecord struct{
	Count string
	Year string
}

func RegisterDB(){
	/*if !com.IsExist(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}*/
	//注册模型
	orm.RegisterModel(new(Category),new(Topic),new(Comment))
	orm.RegisterDataBase("default","mysql","root:200934652qwe@tcp(127.0.0.1:3306)/charset=utf8&parseTime=True&loc=Local",30)
	db=orm.NewOrm()
}
func CheckUser(username string,password string) bool{
	if username=="admin" && password=="123456"{
		return true
	}
	return false


}
func AddCategory(category *Category) bool{
	_,err:=db.Insert(category)
	if err!=nil{
		fmt.Println("err:",err)
		db.Rollback()
		return false
	}else{
		return true
	}
}
func ReadAllCategory() ([]Category,error){


    var category []Category
	_, err := db.QueryTable("category").RelatedSel().All(&category)
	if err == nil {
		//fmt.Printf("%d category read\n", num)
		for i, c := range category {
			var topics []Topic
			num, err := db.QueryTable("topic").Filter("Category",c.Id).RelatedSel().All(&topics)
			if err!=nil{
				log.Fatal(err)
			}
			//fmt.Println("num:",num)
			category[i].TopicCount=num

			//fmt.Printf("Id: %d, UserName: %d, Title: %s\n", c.Id,c.User.Id,c.User.Username)
		}
	}
	return category,err
}
func DeleteCategory(category Category)bool{
	_,err:=db.Delete(&category)
	if err!=nil{
		db.Rollback()
		return false
	}
	return true
}

func FindCategoryById(id int) Category{
	qb,err:=orm.NewQueryBuilder("mysql")
	if err!=nil{
		log.Fatal(err)
	}
	qb.Select("*").From("category").Where("id=?")
	sql:=qb.String()
	var category Category
	err= db.Raw(sql,id).QueryRow(&category)
	if err!=nil{
		log.Fatal(err)
	}
	return category
}
func AddTopic(topic *Topic) bool{
	_,err:=db.Insert(topic)
	if err!=nil{
		fmt.Println(err)
		db.Rollback()
		return false
	}else{
		return true
	}
}
func ReadAlltopic(isDesc bool,topic []Topic) ([]Topic,error){
	qb,err:=orm.NewQueryBuilder("mysql")
	if err!=nil{
		log.Fatal(err)
	}
	if isDesc{
		qb.Select("*").From("topic").OrderBy("created").Desc()
	}else{
		qb.Select("*").From("topic")
	}
	sql:=qb.String()
	_,err=db.Raw(sql).QueryRows(&topic)
	//关联查询读取每篇文章评论数
	for i,_:=range topic{
		var comments Comment
		num1,_:=db.QueryTable("comment").Filter("Topic",topic[i].Id).RelatedSel().All(&comments)
		topic[i].ReplyCount=num1

	}
	return topic,err
}
func ReadAllCategories(category *[]Category) (int64,error){
	qb,err:=orm.NewQueryBuilder("mysql")
	if err!=nil{
		log.Fatal(err)
	}
	qb.Select("*").From("category")

	sql:=qb.String()
	return db.Raw(sql).QueryRows(category)
}
//分页
func ReadTopicByPage(pageNow int, pageSize int) ([]Topic, error){
	sql:="select * from topic  limit ?,?"
	var topics []Topic
	_,err:=db.Raw(sql,(pageNow-1)*pageSize,pageSize).QueryRows(&topics)

	return topics,err
}
//获取pageCount
func GetTopicPageCount(pageSize int) int{
	var rowCount int
	sql:="select count(*) from topic "
	err:=db.Raw(sql).QueryRow(&rowCount)
	if err!=nil{
		log.Fatal(err)
	}
	if rowCount%pageSize==0{
		return rowCount/pageSize
	}else{
		return (rowCount/pageSize)+1
	}
}
func DeleteTopic(topic Topic)bool{
	_,err:=db.Delete(&topic)
	if err!=nil{
		db.Rollback()
		return false
	}
	return true
}
func FindTopicById(id int) Topic{
	qb,err:=orm.NewQueryBuilder("mysql")
	if err!=nil{
		log.Fatal(err)
	}
	qb.Select("*").From("topic").Where("id=?")
	sql:=qb.String()
	var topic Topic
	err= db.Raw(sql,id).QueryRow(&topic)
	if err!=nil{
		log.Fatal(err)
	}
	return topic
}
func ReadOneTopic(topic *Topic) error{
	qb,err:=orm.NewQueryBuilder("mysql")
	if err!=nil{
		log.Fatal(err)
	}
	qb.Select("*").From("topic")
	sql:=qb.String()
	return db.Raw(sql).QueryRow(topic)
}
func UpdateTopic(topic *Topic) bool{
	_,err:=db.Update(topic)


	if err!=nil{
		fmt.Println(err)
		db.Rollback()
		return false
	}else{
		return true
	}
}
func AddComment(comment *Comment) bool{

	topic:=FindTopicById(int(comment.Topic.Id))
	topic.ReplyCount++
	UpdateTopic(&topic)

	_,err:=db.Insert(comment)
	if err!=nil{
		fmt.Println("err:",err)
		db.Rollback()
		return false
	}else{
		return true
	}
}
func ReadTopicComments(topicId int) ([]Comment,error){
	qb,err:=orm.NewQueryBuilder("mysql")
	if err!=nil{
		log.Fatal(err)
	}
    var comments []Comment

	qb.Select("*").From("comment").Where("topic_id=?")

	sql:=qb.String()
	_,err=db.Raw(sql,topicId).QueryRows(&comments)
	//fmt.Println("comments:",comments)
	return comments,err
}
func ReadAllTopicByCategory() ([]Category){
	var categories []Category
	_,_=db.QueryTable("category").All(&categories)

	for i,_:=range categories{
		var topics []Topic
		num,_:=db.QueryTable("topic").Filter("Category",categories[i].Id).RelatedSel().All(&topics)
		categories[i].TopicCount=num
		//fmt.Println(categories[i].Title,categories[i].TopicCount)
	}
	return categories


}
//分页
func ReadCategoryTopicByPage(categoryId int,pageNow int, pageSize int) ([]Topic, error){
	sql:="select * from topic where category_id=? limit ?,?"
	var topics []Topic
	_,err:=db.Raw(sql,categoryId,(pageNow-1)*pageSize,pageSize).QueryRows(&topics)

	return topics,err
}
//获取pageCount
func GetCategoryTopicPageCount(categoryId int,pageSize int) int{
	var rowCount int
	sql:="select count(*) from topic where category_id=? "
	err:=db.Raw(sql,categoryId).QueryRow(&rowCount)
	if err!=nil{
		log.Fatal(err)
	}
	if rowCount%pageSize==0{
		return rowCount/pageSize
	}else{
		return (rowCount/pageSize)+1
	}
}

//查询所有月份文章数目
func ReadAllTopicsByTime() ([]TopicRecord){
	sql:="SELECT count(*) as count,DATE_FORMAT(created,'%Y-%m') as year FROM topic group by DATE_FORMAT(created,'%Y-%m')"
	var m []orm.Params
	num,err:=db.Raw(sql).Values(&m)
	if err!=nil{
		fmt.Println(err)
		return nil
	}
	var topicRecords []TopicRecord

	topicRecords=make([]TopicRecord,num)
	for i,_:=range m{
		//fmt.Println("i:",i)
		topicRecords[i].Count=m[i]["count"].(string)
		topicRecords[i].Year=m[i]["year"].(string)
		//fmt.Println(reflect.TypeOf(m[i]["count"]),m[i]["count"],m[i]["year"])
	}
	//fmt.Println("topicRecord:",topicRecords)
	return topicRecords
}

//分页
func ReadCategoryTopicByMonth(month string,pageNow int, pageSize int) ([]Topic, error){
	sql:="SELECT * FROM topic where DATE_FORMAT(created,'%Y-%m')=? limit ?,?"
	var topics []Topic
	_,err:=db.Raw(sql,month,(pageNow-1)*pageSize,pageSize).QueryRows(&topics)

	return topics,err
}
//获取pageCount
func GetMonthTopicPageCount(month string,pageSize int) int{
	var rowCount int
	sql:="select count(*) from topic where DATE_FORMAT(created,'%Y-%m')=? "
	err:=db.Raw(sql,month).QueryRow(&rowCount)
	if err!=nil{
		log.Fatal(err)
	}
	if rowCount%pageSize==0{
		return rowCount/pageSize
	}else{
		return (rowCount/pageSize)+1
	}
}