package admin

import (
	common "github.com/huanzz/bgadmin_example/common"
	admin "github.com/huanzz/bgadmin_example/models/admin"
	m "github.com/huanzz/bgadmin_example/models/article"
	"log"
	"strconv"
)

type ArticleController struct {
	BaseController
}

//--------------- 文章管理 ----------------------------
func (this *ArticleController)ArticleList()  {
	search := this.GetString("search")
	page,_ := this.GetInt("page")
	if page == 0{
		page = 1
	}
	articles, _ := m.GetArticleList(page,10,search,this.Member.Id)
	total := m.TotalArticleList(search,this.Member.Id)
	this.Data["articles"] = articles
	this.Data["search"] = search
	this.Data["paginator"] = common.Paginator(page,10, total)
	this.TplName = "admin/article/article.html"
}

func (this *ArticleController)ArticleAddPage()  {
	category := m.Category{Id:0}
	article := m.Article{Status:1, Member:&this.Member, Category:&category}
	cateSelect,_ := m.GetCategoryListByPid(0)
	tagList := m.GetTagSelect()
	var tagSelect string
	for _,v := range tagList {
		tagSelect = tagSelect + "<input type='checkbox' class='icheckbox_flat-grey' name='Tags' value='"+
			strconv.Itoa(v.Id)+"' >"+v.Name+"<label>&nbsp;&nbsp;&nbsp;</label>"
	}
	this.Data["article"] = article
	this.Data["cateSelect"] = cateSelect
	this.Data["tagSelect"] = tagSelect
	this.TplName = "admin/article/articleEdit.html"
}

func (this *ArticleController)ArticleEditPage()  {
	Id,_ := this.GetInt("Id")
	cateSelect,_ := m.GetCategoryListByPid(0)
	tagList := m.GetTagSelect()
	article := m.GetArticleById(Id)
	articleTag := m.GetTagListByArticleId(Id)
	var tagSelect string
	var Tid int
	for _,v := range tagList {
		for _,val := range articleTag{
			if v.Id == val.Id {
				tagSelect = tagSelect + "<input type='checkbox' class='icheckbox_flat-grey' name='Tags' value='"+
					strconv.Itoa(v.Id)+"' checked >"+v.Name+"<label>&nbsp;&nbsp;&nbsp;</label>"
				Tid = val.Id
			}
		}
		if v.Id != Tid {
			tagSelect = tagSelect + "<input type='checkbox' class='icheckbox_flat-grey' name='Tags' value='"+
				strconv.Itoa(v.Id)+"' >"+v.Name+"<label>&nbsp;&nbsp;&nbsp;</label>"
		}
	}
	this.Data["article"] = article
	this.Data["cateSelect"] = cateSelect
	this.Data["tagSelect"] = tagSelect
	this.TplName = "admin/article/articleEdit.html"
}

func (this *ArticleController)ArticleSave()  {
	categoryId,_ := this.GetInt("Category")
	MemberId,_ := this.GetInt("Member")
	taglist := this.GetStrings("Tags")
	article := m.Article{}
	if err := this.ParseForm(&article); err != nil {
		this.NoteAndJump("Warning","获取文章信息失败,"+err.Error(),"/admin/article/list")
	}
	var err error
	str := ""
	category := m.GetCategoryById(categoryId)
	member := admin.GetMemberById(MemberId)
	article.Category = &category
	article.Member = &member
	if article.Id > 0{
		_,err = m.UpdateArticle(article)
		m.UpdateTagOfArticle(taglist,article.Id)
		str = "更新"
	} else {
		var Id int64
		Id,err = m.InsertArticle(article)
		articleId,_ := strconv.Atoi(strconv.FormatInt(Id,10))
		m.AddTagToArticle(taglist,articleId)
		str = "添加"
	}

	if err == nil {
		this.NoteAndJump("Success",str+"文章信息成功", "/admin/article/list")
	}else {
		this.NoteAndJump("Error",str+"文章信息失败", "/admin/article/list")
		log.Println(err)
	}
}

func (this *ArticleController)ArticleDel()  {
	Id,_ := this.GetInt("Id")
	if m.DelAllTagFromArticle(Id) {
		_,err := m.DelArticleById(Id)
		if err == nil {
			this.NoteAndJump("Success","删除文章成功", "/admin/article/list")
		}else {
			this.NoteAndJump("Error","删除文章失败", "/admin/article/list")
		}
	}else {
		this.NoteAndJump("Error","清理标签失败，无法删除文章", "/admin/article/list")
	}


}


//--------------- 文章分类管理 ----------------------------
func (this *ArticleController)CategoryList()  {
	Pid,_ := this.GetInt("Pid")
	search := this.GetString("search")
	page,_ := this.GetInt("page")
	var pageSize int
	if page == 0{
		page = 1
	}
	if Pid == 0 {
		pageSize = 10
	}else {
		pageSize = 50
	}
	categories,count := m.GetCategoryList(page, pageSize, search, Pid)
	this.Data["categories"] = categories
	this.Data["search"] = search
	this.Data["ParentCate"] = m.GetCategoryById(Pid)
	this.Data["paginator"] = common.Paginator(page,pageSize, count)
	this.TplName = "admin/article/category.html"
}

func (this *ArticleController)CategoryAddPage()  {
	category := m.Category{Status:1,Pid:0,Sort:1}
	cateSelect := m.GetCategorySelect(0)
	this.Data["category"] = category
	this.Data["cateSelect"] = cateSelect
	this.TplName = "admin/article/categoryEdit.html"
}

func (this *ArticleController)CategoryEditPage()  {
	Id,_ := this.GetInt("Id")
	category := m.GetCategoryById(Id)
	cateSelect:= m.GetCategorySelect(0)
	this.Data["category"] = category
	this.Data["cateSelect"] = cateSelect
	this.TplName = "admin/article/categoryEdit.html"
}

func (this *ArticleController)CategorySave()  {
	category := m.Category{}
	if err := this.ParseForm(&category); err != nil {
		this.NoteAndJump("Warning","获取分类信息失败,"+err.Error(),"/admin/articlecategory/list")
	}
	var err error
	str := ""
	if category.Id > 0{
		_,err = m.UpdateCategory(&category)
		str = "更新"
	} else {
		_,err = m.InsertCategory(&category)
		str = "添加"
	}
	if err == nil {
		this.NoteAndJump("Success",str+"分类信息成功", "/admin/articlecategory/list")
	}else {
		this.NoteAndJump("Error",str+"分类信息失败", "/admin/articlecategory/list")
	}
}

func (this *ArticleController)CategoryDel()  {
	Id,_ :=this.GetInt("Id")
	_,count := m.GetCategoryListByPid(Id)
	if count > 0 {
		this.NoteAndJump("Error","操作失败, 请删除子分类后再试", "/admin/articlecategory/list")
	}else {
		_,err := m.DelCategoryById(Id)
		if err == nil {
			this.NoteAndJump("Success","删除分类信息成功", "/admin/articlecategory/list")
		}else {
			this.NoteAndJump("Error","删除分类失败", "/admin/articlecategory/list")
		}
	}

}

//--------------- 文章标签管理 ----------------------------

func (this *ArticleController)TagList()  {
	search := this.GetString("search")
	page,_ := this.GetInt("page")
	if page == 0{
		page = 1
	}
	tags,count := m.GetTagList(page, 10, search)
	this.Data["tags"] = tags
	this.Data["search"] = search
	this.Data["paginator"] = common.Paginator(page,10, count)
	this.TplName = "admin/article/tag.html"
}

func (this *ArticleController)TagAddPage()  {
	tag := m.Tag{Status:1}
	this.Data["tag"] = tag
	this.TplName = "admin/article/tagEdit.html"
}

func (this *ArticleController)TagEditPage()  {
	Id,_  := this.GetInt("Id")
	tag := m.GetTagById(Id)
	this.Data["tag"] = tag
	this.TplName = "admin/article/tagEdit.html"
}

func (this *ArticleController)TagSave()  {
	tag := m.Tag{}
	if err := this.ParseForm(&tag); err != nil {
		this.NoteAndJump("Warning","获取标签信息失败,"+err.Error(),"/admin/articletag/list")
	}
	var err error
	str := ""
	if tag.Id > 0{
		_,err = m.UpdateTag(&tag)
		str = "更新"
	} else {
		_,err = m.InsertTag(&tag)
		str = "添加"
	}
	if err == nil {
		this.NoteAndJump("Success",str+"标签信息成功", "/admin/articletag/list")
	}else {
		this.NoteAndJump("Error",str+"标签信息失败", "/admin/articletag/list")
	}
}

func (this *ArticleController)TagDel()  {
	Id,_ :=this.GetInt("Id")
	_,err := m.DelTagById(Id)
	if err == nil {
		this.NoteAndJump("Success","删除标签信息成功", "/admin/articletag/list")
	}else {
		this.NoteAndJump("Error","删除标签失败", "/admin/articletag/list")
	}
}


//--------------- 文章评论 ----------------------------
func (this *ArticleController)CommentList()  {
	articleId,_ := this.GetInt("Id")
	page,_ := this.GetInt("page")
	if page == 0{
		page = 1
	}
	article := m.GetArticleById(articleId)
	comments,_ := m.GetCommentList(page, 100, articleId)
	total := m.TotalComment(articleId)
	this.Data["article"] = article
	this.Data["comments"] = comments
	this.Data["paginator"] = common.Paginator(page,100, total)
	this.TplName = "admin/article/comment.html"
}

func (this *ArticleController)CommentDel()  {
	Id,_ := this.GetInt("Id")
	ArticleId:= this.GetString("ArticleId")
	_,err := m.DelCommentById(Id)
	if err == nil {
		this.NoteAndJump("Success","删除评论成功", "/admin/articlecomment/list?Id="+ArticleId)
	}else {
		this.NoteAndJump("Error","删除评论失败", "/admin/articlecomment/list?Id="+ArticleId)
	}
}



