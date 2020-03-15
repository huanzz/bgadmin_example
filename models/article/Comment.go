package article

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
	"log"
	"errors"
	. "bgadmin_example/common"
)

type Comment struct {
	Id  		int
	Writer 		string			`valid:"Required"`
	Email 		string			`valid:"Email;Required"`
	Context 	string			`valid:"Required"`
	Floor		int				`valid:"Required"`
	Reply		int				`valid:"Required"`
	Article		*Article		`orm:"rel(fk)"`
	Createtime	time.Time		`orm:"type(datetime);auto_now_add"`
}


func init () {
	orm.RegisterModel(new(Comment))
}

// Check Comment
func CheckComment(comment *Comment) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&comment)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

// Insert Comment
func InsertComment(comment Comment) (int64, error) {
	if err := CheckComment(&comment); err != nil{
		return 0, err
	}
	o := orm.NewOrm()
	comm := new(Comment)
	comm.Writer = comment.Writer
	comm.Email = comment.Email
	comm.Context = comment.Context
	comm.Article = comment.Article
	comm.Floor = comment.Floor
	comm.Reply = comment.Reply
	id, err := o.Insert(comm)
	return id,err
}

// Update Comment
func UpdateComment(comment Comment) (int64, error) {
	if err := CheckComment(&comment); err != nil{
		return 0, err
	}
	o := orm.NewOrm()
	comm := Comment{Id: comment.Id}
	if len(comment.Writer)>0{
		comm.Writer = comment.Writer
	}
	if len(comment.Email)>0{
		comm.Email= comment.Email
	}
	if len(comment.Context)>0{
		comm.Context = comment.Context
	}
	if comment.Floor != -1 {
		comm.Floor = comment.Floor
	}
	if comment.Reply != -1 {
		comm.Reply = comment.Reply
	}
	if comment.Article != nil{
		article := GetArticleById(comment.Article.Id)
		comm.Article = &article
	}
	num,err :=o.Update(&comm)
	return num, err
}

// Delete Comment ById
func DelCommentById(id int) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Delete(&Comment{Id: id})
	return num, err
}

// Get Comment List
func GetCommentList(page int, pageSize int, articleId int) (comments []Comment, count int64)  {
	o := orm.NewOrm()
	db := o.QueryTable(new(Comment))
	offset := PageOffset(page, pageSize)
	count,_ = db.Limit(pageSize, offset).Filter("Article__Id", articleId).OrderBy("-Createtime").All(&comments)
	return comments,count
}

// Total Comment
func TotalComment(articleId int) (total int64) {
	o := orm.NewOrm()
	db := o.QueryTable(new(Comment))
	total,_ = db.Filter("Article__Id", articleId).Count()
	return total
}

// comment show in content web
func GetCommentByArticle(articleId int) (comments []Comment ,count int64)  {
	o := orm.NewOrm()
	db := o.QueryTable(new(Comment))
	count,_ = db.Filter("Article__Id",articleId).OrderBy("-Createtime").All(&comments)
	return comments,count
}

// 上一条 Comment Id
func LastComment(articleId int) int {
	var comment Comment
	o := orm.NewOrm()
	db := o.QueryTable(new(Comment))
	db.Filter("Article__Id",articleId).OrderBy("-Floor").One(&comment)
	return comment.Floor
}