package controller

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "guishi_server/models"
	"time"
)

type ShopperController struct {
	beego.Controller
}

func (this *ShopperController) Get() {
	this.TplName = "shopper_test.html"
}

func (this *ShopperController) Post() {

	options := this.Input().Get("options")
	beego.Info(options)
	//检查请求的方法
	if options != "" {
		switch options {
		//登录
		case "login":
			this.login()
		//注册
		case "register":
			this.register()
		case "createStore":
			this.createStore()
		case "updateStore":
			this.updateStore()
		case "addProductType":
			this.addProductType()
		case "addGoods":
			this.addGoods()
		default:
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}
	this.Data["json"] = map[string]interface{}{"status": 400, "msg": "options is null !", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}
