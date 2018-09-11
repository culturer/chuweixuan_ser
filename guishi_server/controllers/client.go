package controller

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "guishi_server/models"
	"time"
)

type ClientController struct {
	beego.Controller
}

func (this *ClientController) Get() {
	this.TplName = "client_test.html"
}

func (this *ClientController) Post() {

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
		//获取首页广告轮播
		case "getBanner":
			this.getBanner()
		//获取商铺
		case "getShopper":
			this.getShopper()
			//获取商品
		case "getGoods":
			this.getGoods()
			//搜索商品
		case "searchGoods":
			this.searchGoods()
			//获取订单
		case "getOrders":
			this.getOrders()
			//提交订单
		case "submitOrder":
			this.submitOrder()
			//发货订单
		case "sendOrder":
			this.sendOrder()
			//签收订单
		case "receiveOrder":
			this.receiveOrder()
			//支付订单
		case "payOrder":
			this.payOrder()
			//取消订单
		case "cancelOrder":
			this.cancelOrder()
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
