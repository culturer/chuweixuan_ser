package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//用户表
type TUser struct {
	//用户基本信息
	Id int64
	// 用户名
	Tel string
	// 密码
	Pwd string
	//收款人
	Payee string
	//地址
	Address string
	// 收款账号
	Amount string
	// 账号类别
	AmountType string
	//是否是消费者
	IsCustomer bool
	//是否是商家
	IsSeller bool
	//是否是配送员
	IsDiliver bool
	//是否是管理员
	IsManager bool
	//微信openId
	Vid string
	//是否冻结
	IsLock bool
	//创建时间 --- 时间戳
	AddTime int64
}

//新建用户
func AddUser(user *TUser) (int64, error) {
	o := orm.NewOrm()
	userId, err := o.Insert(user)
	return userId, err
}

//查询账号
func GetUserById(userId int64) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser)
	qs := o.QueryTable("t_user")
	err := qs.Filter("id", userId).One(user)
	return user, err
}

//手机号查询账号
func GetUserByTel(tel string) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser)
	qs := o.QueryTable("t_user")
	err := qs.Filter("tel", tel).One(user)
	return user, err
}

//微信Id查询账号
func GetUserByVId(vid string) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser)
	qs := o.QueryTable("t_user")
	err := qs.Filter("vid", vid).One(user)
	return user, err
}
