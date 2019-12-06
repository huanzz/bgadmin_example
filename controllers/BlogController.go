package controllers

import (
	common "bgadmin/common"
	model "bgadmin/models/article"
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
	search string
}

func (this *BlogController)Prepare() {
	this.search = this.GetString("search")
	categories,_ := model.GetCategoryListByPid(0)
	this.Data["search"] = this.search
	this.Data["publishArticles"] = model.GetArticlesByPublish()
	this.Data["sidebarTags"] = model.GetTagSelect()
	this.Data["sidebarCategory"] = categories
	this.Layout = "blog/layout.html"
}


func (this *BlogController)Index()  {
	page,_ := this.GetInt("page")
	if page == 0{
		page = 1
	}
	articles,_ := model.ArticlesByTime(page, 10, this.search)
	total := model.TotalArticlesByTime(this.search)
	this.Data["articles"] = articles
	this.Data["articleBy"] = "最新文章"
	this.Data["paginator"] = common.Paginator(page,10, total)
	this.Data["total"] = total
	this.TplName = "blog/index.html"
}

func (this *BlogController)Category()  {
	category := this.GetString("name")
	page,_ := this.GetInt("page")
	if page == 0{
		page = 1
	}
	articles,_ := model.ArticlesByTagOrCategory(page, 10,"category", category)
	total := model.TotalArticlesByTagOrCategory("category", category)
	this.Data["articles"] = articles
	this.Data["articleBy"] = "类别：" + category
	this.Data["paginator"] = common.Paginator(page,10, total)
	this.Data["total"] = total
	this.TplName = "blog/index.html"
}

func (this *BlogController)Tag()  {
	tag := this.GetString("name")
	page,_ := this.GetInt("page")
	if page == 0{
		page = 1
	}
	articles,_ := model.ArticlesByTagOrCategory(page, 10,"tag", tag)
	total := model.TotalArticlesByTagOrCategory("tag", tag)
	this.Data["articles"] = articles
	this.Data["articleBy"] = "标签：" + tag
	this.Data["paginator"] = common.Paginator(page,10, total)
	this.Data["total"] = total
	this.TplName = "blog/index.html"
}

func (this *BlogController)Content()  {
	articleId,_:= this.GetInt("Id")
	article,err := model.GetOneArticle(articleId)
	if err != nil{
		this.Abort("404")
	}
	commemts,_ := model.GetCommentByArticle(articleId)
	preArticle, err := model.GetOneArticle(articleId - 1)
	nextArticle, err := model.GetOneArticle(articleId + 1)
	this.Data["article"] = article//上一篇
	this.Data["pre_article"] = preArticle//上一篇
	this.Data["next_article"] = nextArticle//下一篇
	this.Data["commemts"] = commemts
	this.TplName = "blog/content.html"
}

func (this *BlogController)Comment()  {
	writer := this.GetString("writer")
	email := this.GetString("email")
	context := this.GetString("context")
	articleId,_ := this.GetInt("articleId")
	article := model.GetArticleById(articleId)
	curFloor := model.LastComment(articleId)
	comment := model.Comment{Writer:writer,Email:email,Context:context,Article:&article,Floor:curFloor+1}
	_,err := model.InsertComment(comment)
	res := make(map[string]string)
	if err!= nil {
		res["msg"] = err.Error()
	}else {
		res["msg"] ="评论成功"
	}
	this.Data["json"] = res
	this.ServeJSON()
	return
}
