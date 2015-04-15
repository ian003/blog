package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" //"_"表示只执行初始化函数
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/blog.db"
	_SQLITES_DRIVER = "sqlite3"
)

//种类
type Category struct {
	Id              int64
	Title           string    //默认255字节
	Created         time.Time `orm:"index"` //创建索引
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

//文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:size(5000)`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {

	//判断数据库是否存在，不存在则创建
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	//创建ORM
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITES_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITES_DRIVER, _DB_NAME, 10)

}

//添加分类
func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}

	qs := o.QueryTable("category")

	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err1 := o.Insert(cate)

	if err1 != nil {
		return err1
	}

	return nil
}

//删除分类
func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

//获取所有分类
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	_, err := qs.All(&cates)

	return cates, err
}

//添加文章
func AddTopic(title string, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}
	_, err := o.Insert(topic)
	return err
}

//获取所有文章
func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics) //倒序,以数据库中字段为准
	} else {
		_, err = qs.All(&topics)

	}

	return topics, err
}

//获取文章
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	topic.Updated = time.Now()
	_, err = o.Update(topic)
	return topic, nil
}

//修改文章
func ModifyTopic(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}

	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return nil
}

//删除文章
func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	return err
}
