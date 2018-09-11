package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"guishi_server/models"
	"time"
)

type Orders struct {
	OrderType string
	OrderData []*models.TOrder
	Count     int
}

//获取订单列表
func (this *ClientController) getOrders() {

	userId, err := this.GetInt64("userId")
	beego.Info(userId)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

	size, err := this.GetInt("size")
	beego.Info(size)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

	index, err := this.GetInt("index")
	beego.Info(index)
	if index == 0 {
		//index从1开始计数
		index = 1
	}
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

	types := []string{"is_pay", "is_send", "is_receive"}
	datas := make([]Orders, 0)
	where := ""
	for i := 0; i < len(types); i++ {
		if i == 0 {
			where = fmt.Sprintf("where user_id = %v and %v = 0", userId, types[i])

		} else if i == 1 {
			where = fmt.Sprintf("where user_id = %v and is_pay = 1 and %v = 0", userId, types[i])
		} else {
			where = fmt.Sprintf("where user_id = %v and is_pay = 1 and is_send = 1 and %v = 0", userId, types[i])

		}
		// where := fmt.Sprintf("where user_id = %v ", userId)
		items, count, err := models.GetOrderPage(index, size, where)
		if err != nil {
			beego.Info(err)
		}
		data := Orders{
			OrderType: types[i],
			OrderData: items,
			Count:     count,
		}
		datas = append(datas, data)
	}

	this.Data["json"] = map[string]interface{}{"status": 200, "datas": datas, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

//提交订单
func (this *ClientController) submitOrder() {
	order := &models.TOrder{}
	err := this.ParseForm(order)
	order.AddTime = time.Now().Unix()
	beego.Info(order)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	order.AddTime = time.Now().Unix()
	order.IsPay = false
	orderId, err := models.AddOrder(order)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "orderId": orderId, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

//订单发货
func (this *ClientController) sendOrder() {
	orderId, err := this.GetInt64("orderId")
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	order, err := models.GetOrderById(orderId)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	order.IsSend = true
	sendType, err := this.GetInt("sendType")
	order.SendType = sendType
	err = models.UpdateOrder(order)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}

	//在此加入发货逻辑，比如发送送货通知，推送送货单信息给快递员接单

	this.Data["json"] = map[string]interface{}{"status": 200, "order": order, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
}

//订单签收
func (this *ClientController) receiveOrder() {
	orderId, err := this.GetInt64("orderId")
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	order, err := models.GetOrderById(orderId)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	if order.IsSend == true && order.IsPay == true && order.IsCancel == false {
		order.IsReceive = true
		order.ReceiveTiem = time.Now().Unix()
	}
	err = models.UpdateOrder(order)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "orderId": orderId, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
}

//订单取消
func (this *ClientController) cancelOrder() {
	orderId, err := this.GetInt64("orderId")
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	order, err := models.GetOrderById(orderId)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	//客户已经付款时取消订单
	if order.IsPay == true && order.IsSend == false {

	}
	//商家已经派送时取消订单
	if order.IsSend == true && order.IsPay == true {

	}
	if order.IsSend == true && order.IsPay == false {

	}
	//配送员已经接单时取消订单
	if order.IsSend1 == true {

	}
	//客户已经签收时取消订单
	if order.IsReceive == true {

	}
}

func (this *ClientController) payOrder() {

	orderId, err := this.GetInt64("orderId")
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	order, err := models.GetOrderById(orderId)
	if err != nil {
		beego.Info(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
	}
	payType := this.GetString("payType")
	order.PayType = payType
	//执行支付逻辑
	if payType != "" {
		switch payType {
		//微信支付
		case "weixin":
			this.weixinPay()
		//支付宝支付
		case "zhifubao":
			this.zhifubaoPay()
		// 现金支付
		case "cash":
			this.cashPay()
		default:
			this.Data["json"] = map[string]interface{}{"status": 400, "msg": "没有对应处理方法", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
	}
	this.Data["json"] = map[string]interface{}{"status": 400, "msg": "账号或密码不为空 ，请检查后重新登录！", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

func (this *ClientController) weixinPay() error {

	return nil
}

func (this *ClientController) zhifubaoPay() error {

	return nil
}

func (this *ClientController) cashPay() error {

	return nil
}
