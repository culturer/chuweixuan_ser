package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//订单表
type TOrder struct {
	//订单编号
	Id int64
	//用户编号
	UserId int64
	// 联系电话
	Tel string
	// 联系地址
	Address string
	// 收货人
	Receiver string
	// 金额
	Amount float64
	// 付款类型
	PayType string
	// 是否折扣
	IsDiscounts bool
	// 打折信息 --- JSON保存
	DiscountMsg string
	// 备注
	Msg string
	// 商品项 --- JSON保存
	Item string `orm:"type(text);null"`
	// 是否支付
	IsPay bool
	// 是否预约
	Appoint bool
	// 预约时间
	AppointTime int64
	// 是否发货
	IsSend bool
	//快递员接单
	IsSend1 bool
	// 派送员编号
	SenderId int64
	// 发货方式
	SendType int
	// 是否收货
	IsReceive bool
	// 收货时间 --- 时间戳
	ReceiveTiem int64
	// 取消订单
	IsCancel bool
	// 退单理由
	CancelMsg string
	// 创建时间 --- 时间戳
	AddTime int64
}

//添加订单
func AddOrder(order *TOrder) (int64, error) {
	o := orm.NewOrm()
	orderId, err := o.Insert(order)
	return orderId, err
}

//修改订单
func UpdateOrder(order *TOrder) error {
	o := orm.NewOrm()
	_, err := o.Update(order)
	return err
}

//查询账号
func GetOrderById(orderId int64) (*TOrder, error) {
	o := orm.NewOrm()
	order := new(TOrder)
	qs := o.QueryTable("t_order")
	err := qs.Filter("id", orderId).One(order)
	return order, err
}

//分页获取数据
func GetOrderPage(index, size int, where string) ([]*TOrder, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TOrder{}
	dataCounts := []*TOrder{}
	//返回数据列表
	sql := fmt.Sprintf("select * from t_order %v  order by id desc limit %v offset %v", where, size, size*(index-1))
	beego.Info(sql)
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数
	sqlCount := fmt.Sprintf("select * from t_order  %v ", where)
	count, err1 := ormHelper.Raw(sqlCount).QueryRows(&dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	beego.Info(count)
	return data, int(count), err
}
