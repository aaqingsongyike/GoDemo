package models

import(
	"fmt"
	"strings"
	"strconv"
	"os"
	"time"
	"path"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_"github.com/mattn/go-sqlite3"
)

const(
	_DB_NAME ="data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

//分类
type Category struct{	//模型1
	Id int64
	Title string
	Created time.Time `orm:"index"`
	View int64 `orm:"index"`	//浏览次数
	TopicTime time.Time `orm:"index"`	//最后一篇发表的时间
	TopicCount int64
	TopicLastUserId int64
}

//文章
type Topic struct{	//模型2
	Id int64
	Uid int64
	Title string
	Category string
	Labels string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"`
	Update time.Time `orm:"index"`
	View int64 `orm:"index"`	//浏览次数
	Author string
	ReplyTime time.Time `orm:"index"`
	ReplyCount int64
	ReplyLastUserId int64
}

//评论
type Comment struct {
	Id int64
	Tid int64
	Name string
	Content string	`orm:"size(1000)"`
	Created time.Time	`orm:"index"`
}

func RegisterDB(){
	//检查文件是否存在
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)	//取出目录的路径，并设置默认权限
		os.Create(_DB_NAME) //创建文件
	}

	//创建模型
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	//注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	//创建默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {	//添加操作
	o := orm.NewOrm()	//获取orm对象

	cate := &Category{Title: name}	//创建category对象

	//查询判断name是否被用
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)

	if err == nil {	//表示找到
		return err
	}

	//不存在
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

//删除操作
func DelCategory(id string) error {

	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

//获取所有文章操作
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

//添加文章操作
func AddTopic(title, category, label, content, attachment string) error {
	//处理标签
	/*
	空格作为多个标签的分隔符
	设一个标签为beego   存到数据库中为    $beego#
	*/
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"


	o := orm.NewOrm()
	
	topic := &Topic{	//创建Topic对象
		Title: title,
		Category: category,
		Labels: label,
		Content: content,
		Attachment: attachment,
		Created: time.Now(),
		Update: time.Now(),
	}

	_, err := o.Insert(topic)	//执行插入操作 
	if err != nil {
		return err
	}
	//更新分类统计
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		//如果不存在，简单的忽略更新操作
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	return nil
}

//获取所有文章操作
func GetAllTopics(cate, label string, isDesc bool) ([]*Topic, error) {	//参数（是否倒序排序）
	
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	var err error
	if isDesc {		
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)	//倒序排序
	}else{
		_, err = qs.All(&topics)
		fmt.Println(err)
	}
	return topics, err
}

//获取单独的文章
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
	topic.View++
	_, err = o.Update(topic)
	//放在Update之后
	topic.Labels = strings.Replace(strings.Replace(
		topic.Labels, "#", " ", -1), "$", "", -1)
	return topic, err
}

//Modify
//修改操作
func ModifyTopic(tid, title, category, label, content, attachment string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	//处理标签
	/*
	空格作为多个标签的分隔符
	设一个标签为beego   存到数据库中为    $beego#
	*/
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	var oldCate, oldAttach string

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Labels = label
		topic.Content = content
		topic.Attachment = attachment
		topic.Category = category
		topic.Update = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}

	//更新分类统计
	if len(oldCate) > 0 {	//更新旧的
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", category).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	//删除旧的附件
	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}


	return nil
}

//删除操作
func DeleteModify(tid string) error{
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	var oldCate string

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}

	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}

	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	//_, err = o.Delete(topic)
	return err
}

//Reply
//评论
func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Comment{
		Tid: tidNum,
		Name: nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err == nil {
		return err
	}

	topic := &Topic{Id: tidNum}	//获得Topic对象
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}

	return err
}

//获取所有replies
func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies = make([]*Comment, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}

//删除reply
func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	var tidNum int64
	reply := &Comment{Id: ridNum}
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}

	replies := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_,err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {	
		topic.ReplyTime = replies[0].Created
		topic.ReplyCount = int64(len(replies))
		_, err = o.Update(topic)
	}

	return err
}