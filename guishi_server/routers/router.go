package routers

import (
	"github.com/astaxie/beego"
	"guishi_server/controllers"
)

func init() {

	//初始化路由
	initRouter()
	//初始化过滤器
	// initFilter()

}

//初始化路由
func initRouter() {
	//客户端
	beego.Router("/client", &controller.ClientController{})
	beego.Router("/shopper", &controller.ShopperController{})
	beego.Router("/utils", &controller.UtilsController{})

}

// //初始化过滤器
// func initFilter() {
// 	// beego.InsertFilter("/*", beego.BeforeRouter, user_filter)
// 	//登录过滤器
// 	beego.InsertFilter("/*", beego.BeforeRouter, login_filter)
// 	//供应商采购页面过滤器
// 	beego.InsertFilter("/procurement", beego.BeforeRouter, p_login_filter)
// }

// // func user_filter(ctx *context.Context) {
// // 	//过滤的url表
// // 	fileter_url := []string{"/login", "/register", "/products", "/get", "/wxhelper"}
// // 	//pid --- 分销商Id
// // 	for i := 0; i < len(fileter_url); i++ {
// // 		if ctx.Request.RequestURI == fileter_url[i] {
// // 			return
// // 		}
// // 	}
// // 	_, ok := ctx.Input.Session("uid").(int64)
// // 	if !ok {
// // 		//beego.Info(fmt.Sprintf("redirect,uid:%v", uid))
// // 		//ctx.Redirect(302, "/login")
// // 		ctx.Output.Body([]byte(`{"status":"302","msg":"请重新登陆"}`))
// // 	}
// // }

// //登录过滤器
// func login_filter(ctx *context.Context) {

// 	//不过滤的url表
// 	n_fileter_url := []string{"/partner", "/login", "/register", "/products", "/advertise", "/get", "/procurement", "/p_login", "/wxhelper"}

// 	for i := 0; i < len(n_fileter_url); i++ {
// 		if ctx.Request.RequestURI == n_fileter_url[i] {
// 			return
// 		}
// 	}
// 	//uid --- 用户id
// 	_, ok := ctx.Input.Session("uid").(int64)
// 	if !ok {
// 		ctx.Output.Body([]byte(`{"status":"302","msg":"请重新登陆"}`))
// 	}

// }

// //gi
// func p_login_filter(ctx *context.Context) {

// 	//过滤的url表
// 	fileter_url := []string{"/procurement"}
// 	//pid --- 分销商Id
// 	for i := 0; i < len(fileter_url); i++ {
// 		if ctx.Request.RequestURI == fileter_url[i] {
// 			_, ok := ctx.Input.Session("pid").(int64)
// 			if !ok {
// 				beego.Info("partnerId is null")
// 				//跳转到分销商登录页面
// 				ctx.Redirect(302, "/p_login")
// 			}
// 		}
// 	}

// }
