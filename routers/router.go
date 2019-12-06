package routers

import (
	"github.com/huanzz/bgadmin_example/controllers"
	"github.com/huanzz/bgadmin_example/controllers/admin"
	"github.com/huanzz/bgadmin_example/filter"
	"github.com/astaxie/beego"
)

func init() {

	//Error 页面
	beego.ErrorController(&controllers.ErrorControl{})

	//Blog
	beego.Router("/blog/index", &controllers.BlogController{}, "*:Index")
	beego.Router("/blog/category", &controllers.BlogController{}, "*:Category")
	beego.Router("/blog/tag", &controllers.BlogController{}, "*:Tag")
	beego.Router("/blog/content", &controllers.BlogController{}, "*:Content")
	beego.Router("/blog/comment", &controllers.BlogController{}, "*:Comment")


	//前台及登录相关
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/index", &controllers.IndexController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{},"*:Login")
	beego.Router("/logout", &controllers.LoginController{},"*:Logout")
	beego.Router("/loginadmin", &controllers.LoginController{},"*:LoginAdmin")
	beego.Router("/changepwd", &controllers.LoginController{},"*:ChangePwd")
	beego.Router("/savechange", &controllers.LoginController{},"*:SaveChange")
	beego.Router("/notice", &controllers.LoginController{},"*:Notice")


	//系统后台登录过滤
	beego.InsertFilter("/admin/*",beego.BeforeRouter,filter.FilterAdmin)

	//系统后台-基本操作-
	beego.Router("/admin/operate/tips", &admin.BaseController{},"*:Tips")
	beego.Router("/admin/operate/jump", &admin.BaseController{},"*:NoteAndJump")
	beego.Router("/admin/operate/person", &admin.AdminController{},"*:Person")
	beego.Router("/admin/operate/savemsg", &admin.AdminController{},"*:SaveMsg")

	//系统后台-首页-
	beego.Router("/admin/index", &admin.AdminController{},"*:Index")
	beego.Router("/admin/index/index", &admin.AdminController{},"*:Index")

	//系统后台-会员-
	beego.Router("/admin/member/list", &admin.AdminController{},"*:MemberList")
	beego.Router("/admin/member/add", &admin.AdminController{},"*:MemberAddPage")
	beego.Router("/admin/member/edit", &admin.AdminController{},"*:MemberEditPage")
	beego.Router("/admin/member/doedit", &admin.AdminController{},"*:MemberEdit")
	beego.Router("/admin/member/del", &admin.AdminController{},"*:MemberDel")

	//系统后台-权限组-
	beego.Router("/admin/auth/list", &admin.AdminController{},"*:GroupList")
	beego.Router("/admin/auth/add", &admin.AdminController{},"*:GroupAddPage")
	beego.Router("/admin/auth/edit", &admin.AdminController{},"*:GroupEditPage")
	beego.Router("/admin/auth/doedit", &admin.AdminController{},"*:GroupEdit")
	beego.Router("/admin/auth/del", &admin.AdminController{},"*:GroupDel")
	beego.Router("/admin/auth/authorize", &admin.AdminController{},"*:AuthorizePage")
	beego.Router("/admin/auth/authorizeto", &admin.AdminController{},"*:Authorize")

	//系统后台-菜单-
	beego.Router("/admin/menu/list", &admin.AdminController{},"*:MenuList")
	beego.Router("/admin/menu/add", &admin.AdminController{},"*:MenuAddPage")
	beego.Router("/admin/menu/edit", &admin.AdminController{},"*:MenuEditPage")
	beego.Router("/admin/menu/doedit", &admin.AdminController{},"*:MenuEdit")
	beego.Router("/admin/menu/del", &admin.AdminController{},"*:MenuDel")
	beego.Router("/menu/menu/icons", &admin.AdminController{},"*:MenuIcon")

	//后台--文章管理--文章
	beego.Router("/admin/article/list", &admin.ArticleController{},"*:ArticleList")
	beego.Router("/admin/article/add", &admin.ArticleController{},"*:ArticleAddPage")
	beego.Router("/admin/article/edit", &admin.ArticleController{},"*:ArticleEditPage")
	beego.Router("/admin/article/save", &admin.ArticleController{},"*:ArticleSave")
	beego.Router("/admin/article/del", &admin.ArticleController{},"*:ArticleDel")

	//后台--文章管理--分类
	beego.Router("/admin/articlecategory/list", &admin.ArticleController{},"*:CategoryList")
	beego.Router("/admin/articlecategory/add", &admin.ArticleController{},"*:CategoryAddPage")
	beego.Router("/admin/articlecategory/edit", &admin.ArticleController{},"*:CategoryEditPage")
	beego.Router("/admin/articlecategory/save", &admin.ArticleController{},"*:CategorySave")
	beego.Router("/admin/articlecategory/del", &admin.ArticleController{},"*:CategoryDel")

	//后台--文章管理--标签
	beego.Router("/admin/articletag/list", &admin.ArticleController{},"*:TagList")
	beego.Router("/admin/articletag/add", &admin.ArticleController{},"*:TagAddPage")
	beego.Router("/admin/articletag/edit", &admin.ArticleController{},"*:TagEditPage")
	beego.Router("/admin/articletag/save", &admin.ArticleController{},"*:TagSave")
	beego.Router("/admin/articletag/del", &admin.ArticleController{},"*:TagDel")

	//后台--文章管理--评论
	beego.Router("/admin/articlecomment/list", &admin.ArticleController{},"*:CommentList")
	beego.Router("/admin/articlecomment/del", &admin.ArticleController{},"*:CommentDel")

}
