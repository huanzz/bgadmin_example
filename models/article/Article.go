package article

import (
	"bgadmin_example/models/admin"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strings"
	"time"
	"log"
	"errors"
	common "bgadmin_example/common"
)

type Article struct {
	Id 			int
	Member		*admin.Member	`orm:"rel(fk)"`
	Title 		string
	Description	string			`orm:"null"`
	Content 	string			`orm:"type(text)"`
	Category	*Category		`orm:"rel(fk)"`
	Tags		[]*Tag			`orm:"rel(m2m)"`
	Read 		int				`orm:"default(0)"`
	Like		int				`orm:"default(0)"`
	Publish		int				`orm:"default(0)"`
	Status		int				`orm:"default(0)"`
	Publishtime time.Time 		`orm:"null;type(datetime)"`
	Createtime	time.Time		`orm:"type(datetime);auto_now_add"`
}

type ArticleShow struct {
	Article
	TagSlice	[]map[string]string
}

func init () {
	orm.RegisterModel(new(Article))
}

func CheckArticle(post *Article) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&post)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func InsertArticle(post Article) (int64, error) {
	if err := CheckArticle(&post); err != nil{
		return 0, err
	}
	o := orm.NewOrm()
	article := new(Article)
	article.Title = post.Title
	article.Description = post.Description
	article.Category = post.Category
	article.Status = post.Status
	article.Publish = post.Publish
	article.Content = post.Content
	article.Member = post.Member
	id, err := o.Insert(article)
	return id,err
}

func UpdateArticle(post Article) (int64, error) {
	if err := CheckArticle(&post); err != nil{
		return 0, err
	}
	o := orm.NewOrm()
	article := Article{Id: post.Id}
	if len(post.Title)>0{
		article.Title = post.Title
	}
	if len(post.Description)>0{
		article.Description= post.Description
	}
	if len(post.Content)>0{
		article.Content = post.Content
	}
	if post.Publish != -1{
		article.Publish = post.Publish
	}
	if post.Status != -1{
		article.Status = post.Status
	}
	if post.Member != nil{
		member := admin.GetMemberById(post.Member.Id)
		article.Member = &member
	}

	if post.Category != nil{
		cate := GetCategoryById(post.Category.Id)
		article.Category = &cate
	}
	num,err :=o.Update(&article)
	return num, err
}

func DelArticleById(id int) (int64, error) {
	delTag := DelAllTagFromArticle(id)
	if delTag {
		o := orm.NewOrm()
		num, err := o.Delete(&Article{Id: id})
		return num, err
	}
	return 0, nil
}

func GetArticleList(page int, pageSize int, search string, memberId int) (articleShow []ArticleShow, count int64)  {
	offset := common.PageOffset(page, pageSize)
	var articles []Article
	o := orm.NewOrm()
	db := o.QueryTable(new(Article))
	str :=  strings.Trim(search, " ")
	if memberId == 1 || memberId == 2 {
		count,_ = db.Filter("Title__icontains",str).Limit(pageSize, offset).OrderBy("Createtime").RelatedSel().All(&articles)
	}else {
		count,_ = db.Filter("Title__icontains",str).Filter("MemberId", memberId).Limit(pageSize, offset).OrderBy("Createtime").RelatedSel().All(&articles)
	}
	for _,v := range articles {
		var tags []Tag
		o.QueryTable(new(Tag)).Filter("Article__Article__Id", v.Id).All(&tags)
		var tagMap []map[string]string
		for _,val := range tags {
			tagMap = append(tagMap, map[string]string{"name": val.Name})
		}
		row := ArticleShow{v, tagMap}
		articleShow  =  append(articleShow, row)
	}
	return articleShow, count
}

func TotalArticleList(search string, memberId int) (count int64) {
	var articles []Article
	o := orm.NewOrm()
	db := o.QueryTable(new(Article))
	str :=  strings.Trim(search, " ")
	if memberId == 1 || memberId == 2 {
		count,_ = db.Filter("Title__icontains",str).OrderBy("Createtime").All(&articles)
	}else {
		count,_ = db.Filter("Title__icontains",str).Filter("MemberId", memberId).All(&articles)
	}
	return count
}

func GetArticleById(id int) (article Article) {
	article = Article{Id: id}
	o := orm.NewOrm()
	o.Read(&article, "Id")
	return article
}

func ArticlesByTime(page int, pageSize int, search string) (articleShow []ArticleShow, count int64) {
	offset := common.PageOffset(page, pageSize)
	var articles []Article
	o := orm.NewOrm()
	db := o.QueryTable(new(Article))
	str :=  strings.Trim(search, " ")
	count,_ = db.Filter("Status", 1).Filter("Title__icontains",str).Limit(pageSize, offset).OrderBy("-Createtime").RelatedSel().All(&articles)
	for _,v := range articles {
		var tags []Tag
		o.QueryTable(new(Tag)).Filter("Article__Article__Id", v.Id).All(&tags)
		var tagMap []map[string]string
		for _,val := range tags {
			tagMap = append(tagMap, map[string]string{"name": val.Name})
		}
		row := ArticleShow{v, tagMap}
		articleShow  =  append(articleShow, row)
	}
	return articleShow, count
}

func TotalArticlesByTime(search string) (count int64) {
	var articles []Article
	o := orm.NewOrm()
	db := o.QueryTable(new(Article))
	str :=  strings.Trim(search, " ")
	count,_ = db.Filter("Status", 1).Filter("Title__icontains",str).All(&articles)
	return count
}

func GetOneArticle(id int) (articleShow ArticleShow, err error) {
	var article Article
	o := orm.NewOrm()
	db := o.QueryTable(new(Article))
	err = db.Filter("Id", id).RelatedSel().One(&article)
	var tags []*Tag
	o.QueryTable(new(Tag)).Filter("Status", 1).Filter("Article__Article__Id",id).All(&tags)
	var tagMap []map[string]string
	for _,tt := range tags{
		tagMap = append(tagMap, map[string]string{"name": tt.Name})
	}
	articleShow = ArticleShow{article, tagMap}
	return articleShow,err
}

func GetArticlesByPublish() (articles []*Article) {
	o:= orm.NewOrm()
	db := o.QueryTable(new(Article))
	db.Filter("Status", 1).Filter("Publish", 1).Limit(8,0).All(&articles)
	return articles
}

func ArticlesByTagOrCategory(page int, pageSize int, key, value string) (articleShow []ArticleShow, count int64) {
	offset := common.PageOffset(page, pageSize)
	var articles []Article
	o := orm.NewOrm()
	db := o.QueryTable(new(Article))
	conn := orm.NewCondition()
	switch key {
	case "category":
		conn = conn.And("Category__Name", value)
	case "tag":
		conn = conn.And("Tags__Tag__Name",value)
	}
	count,_ = db.Filter("Status", 1).SetCond(conn).Limit(pageSize, offset).OrderBy("-Createtime").RelatedSel().All(&articles)
	for _,v := range articles {
		var tags []Tag
		o.QueryTable(new(Tag)).Filter("Article__Article__Id", v.Id).All(&tags)
		var tagMap []map[string]string
		for _,val := range tags {
			tagMap = append(tagMap, map[string]string{"name": val.Name})
		}
		row := ArticleShow{v, tagMap}
		articleShow  =  append(articleShow, row)
	}
	return articleShow, count
}

func TotalArticlesByTagOrCategory( key, value string) (count int64) {
	var articles []Article
	o := orm.NewOrm()
	db := o.QueryTable(new(Article))
	conn := orm.NewCondition()
	str :=  strings.Trim(value, " ")
	switch key {
	case "category":
		conn = conn.And("Category__Name", str)
	case "tag":
		conn = conn.And("Tags__Tag__Name",str)
	}
	count,_ = db.Filter("Status", 1).SetCond(conn).RelatedSel().All(&articles)
	return count
}

