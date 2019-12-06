package article

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
	. "github.com/huanzz/bgadmin_example/common"
)

type Category struct {
	Id 			int
	Pid 		int				`orm:"default(0)" form:"Pid" `
	Name 		string			`orm:"unique" form:"Name" valid:"Required"`
	Types	 	int				`orm:"default(1)"`
	Status 		int				`orm:"default(0)" form:"Status"`
	Sort 		int				`orm:"default(0)" form:"Sort"`
	Active 		bool			`orm:"-"`
	Children 	[]*Category		`orm:"-"`
}

type CategorySelect struct {
	Category
	ChildCount 	int64
}

func init () {
	orm.RegisterModel(new(Category))
}


func CheckCategory(cate *Category) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&cate)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func InsertCategory(cate *Category) (int64, error) {
	if err := CheckCategory(cate); err != nil{
		return 0, err
	}
	o := orm.NewOrm()
	category := new(Category)
	category.Name = cate.Name
	category.Status = cate.Status
	category.Pid = cate.Pid
	category.Types = cate.Types
	category.Sort = cate.Sort
	id, err := o.Insert(category)
	return id,err
}

func UpdateCategory(cate *Category) (int64, error) {
	o := orm.NewOrm()
	category := make(orm.Params)
	if len(cate.Name) > 0 {
		category["Name"] = cate.Name
	}
	if cate.Status != -1 {
		category["Status"] = cate.Status
	}
	if cate.Pid != -1 {
		category["Pid"] = cate.Pid
	}
	if cate.Sort != -1 {
		category["Sort"] = cate.Sort
	}
	if cate.Types != -1 {
		category["Types"] = cate.Types
	}
	if len(category) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Category
	num, err := o.QueryTable(table).Filter("Id", cate.Id).Update(category)
	return num, err
}

func DelCategoryById(Id int) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Category{Id: Id})
	return status, err
}

func GetCategoryById(id int)  (cate Category)  {
	cate = Category{Id: id}
	o := orm.NewOrm()
	o.Read(&cate, "Id")
	return cate
}

func GetCategoryListByPid(pid int) (categories []Category, count int64) {
	o := orm.NewOrm()
	db := o.QueryTable(new(Category))
	count,_ = db.Filter("Status", 1).Filter("Pid", pid).All(&categories)
	return categories, count
}

func GetCategoryList(page int, pageSize int, search string, pid int) (cateSelect []CategorySelect, count int64)  {
	o := orm.NewOrm()
	db := o.QueryTable(new(Category))
	offset := PageOffset(page, pageSize)
	var categories []Category
	count,_ = db.Limit(pageSize, offset).Filter("Pid", pid).Filter("Name__icontains", search).OrderBy("Sort").All(&categories)
	for _,v := range categories {
		row := CategorySelect{Category: v}
		_,num := GetCategoryListByPid(v.Id)
		row.ChildCount = num
		cateSelect = append(cateSelect, row)
	}
	count,_ = db.Filter("Pid", pid).Filter("Name__icontains", search).Count()
	return cateSelect, count
}

func GetCategorySelect(pid int) []*Category   {
	var lists []*Category
	o := orm.NewOrm()
	db := o.QueryTable(new(Category))
	db.Filter("Status", 1).OrderBy("Sort").All(&lists)
	data := BuildData(lists)
	res := MakeTree(0,data)
	return res
}

func BuildData(lists []*Category) map[int]map[int] *Category {
	data := make(map[int]map[int]*Category)
	for _,v := range lists {
		id := v.Id
		pid := v.Pid
		if _,ok := data[pid]; !ok {
			data[pid] = make(map[int]*Category)
		}
		data[pid][id] = v
	}
	return data
}

func MakeTree(index int, data map[int]map[int]*Category) []*Category  {
	tmp := make([]*Category, 0)
	for id,item := range data[index]{
		if data[id] != nil {
			item.Children = MakeTree(id,data)
		}
		tmp = append(tmp, item)
	}
	return tmp
}