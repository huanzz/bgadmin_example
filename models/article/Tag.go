package article

import (
	"github.com/astaxie/beego/validation"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"log"
	"errors"
	. "bgadmin/common"
)

type Tag struct {
	Id			int
	Name 		string		`orm:"unique" form:"Name" valid:"Required"`
	Status		int			`orm:"default(1)" form:"Status"`
	Createtime	time.Time	`orm:"type(datetime);auto_now_add" `

}

func init () {
	orm.RegisterModel(new(Tag))
}

func CheckTag(t *Tag) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&t)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func InsertTag(t *Tag) (int64, error) {
	if err := CheckTag(t); err != nil{
		return 0, err
	}
	o := orm.NewOrm()
	tag := new(Tag)
	tag.Name = t.Name
	tag.Status = t.Status

	id, err := o.Insert(tag)
	return id,err
}

func UpdateTag(t *Tag) (int64, error) {
	if err := CheckTag(t); err != nil{
		return 0, err
	}
	o := orm.NewOrm()
	tag := make(orm.Params)
	if len(t.Name) > 0 {
		tag["Name"] = t.Name
	}
	if t.Status != -1 {
		tag["Status"] = t.Status
	}
	if len(tag) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Tag
	num, err := o.QueryTable(table).Filter("Id", t.Id).Update(tag)
	return num, err
}

func DelTagById(Id int) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Delete(&Tag{Id: Id})
	return num, err
}

func GetTagList(page int, pageSize int, search string)  (tags []Tag, count int64){
	o := orm.NewOrm()
	db := o.QueryTable(new(Tag))
	offset := PageOffset(page, pageSize)
	count,_ = db.Limit(pageSize, offset).Filter("Name__icontains", search).OrderBy("-Createtime").All(&tags)
	count,_ = db.Filter("Name__icontains", search).Count()
	return tags, count
}

func GetTagById(Id int) (tag Tag) {
	tag = Tag{Id: Id}
	o := orm.NewOrm()
	o.Read(&tag, "Id")
	return tag
}

func GetTagSelect() (tags []Tag)  {
	o:=orm.NewOrm()
	db := o.QueryTable(new(Tag))
	db.Filter("Status", 1).All(&tags)
	return tags
}

func AddTagToArticle(tags []string, articleId int) (res bool) {
	res = false
	var err error
	o := orm.NewOrm()
	article := Article{Id: articleId}
	m2m := o.QueryM2M(&article, "Tags")
	for _,v := range tags{
		tid,_ := strconv.Atoi(v)
		tag := Tag{Id:tid}
		_, err = m2m.Add(&tag)
	}
	if err == nil{
		res = true
	}
	return res
}

func UpdateTagOfArticle(tags []string, articleId int) (res bool) {
	res = false
	o := orm.NewOrm()
	var err error
	article := Article{Id: articleId}
	m2m := o.QueryM2M(&article,"Tags")
	num,_ :=  m2m.Count()
	if num > 0 {
		m2m.Clear()
	}
	for _,v := range tags{
		tid,_ := strconv.Atoi(v)
		tag :=Tag{Id: tid}
		_,err = m2m.Add(&tag)
	}
	if err == nil {
		res = true
	}
	return res
}

func DelAllTagFromArticle(articleId int) (res bool) {
	res = false
	o := orm.NewOrm()
	article := Article{Id: articleId}
	m2m := o.QueryM2M(&article,"Tags")
	_,err := m2m.Clear()
	if err == nil {
		res = true
	}
	return res
}

func GetTagListByArticleId(articleId int) (tags []*Tag) {
	o := orm.NewOrm()
	db := o.QueryTable(new(Tag))
	db.Filter("Status", 1).Filter("Article__Article__Id", articleId).All(&tags)
	return tags
}
