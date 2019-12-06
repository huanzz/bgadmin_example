package filter

import (
	"github.com/huanzz/bgadmin_example/common"
	"github.com/huanzz/bgadmin_example/controllers"

	"github.com/astaxie/beego/context"
	m "github.com/huanzz/bgadmin_example/models/admin"
)

// 过滤未登录用户
var FilterAdmin = func(ctx *context.Context) {
	webTitle := ""
	tu := ctx.Input.Session("member_info")
	if tu == nil {
		ctx.Redirect(302, "/login")
	}else{
		url := ctx.Request.RequestURI
		if mber, ok := tu.(m.Member); ok {
			 memberId := mber.Id
			if memberId != 1{
				menuMap := m.GetMenuMap(memberId)
				fullUrl := common.GetFullUrl(url)
				webTitle =menuMap[fullUrl]
				if webTitle == "" {
					base := controllers.LoginController{}
					url := base.URLFor("BaseController.Tips","types","Info","msg","没有操作权限","url","/admin/index")
					ctx.Redirect(302, url)
				}
			}

		}
	}





}