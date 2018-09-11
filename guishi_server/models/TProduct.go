package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	// _ "github.com/go-sql-driver/mysql"
)

//商品表
type TProduct struct {
	//商品编号
	Id int64
	// 商品名
	Name string
	//商品分类
	ProductTypeId int64
	// 商品描述
	Desc string
	// 进货价
	APrice float64
	// 售价
	Price string
	// 库存
	Num int
	// 是否下架
	IsLock bool
	// 图标
	Icon string
	// 销量
	SellNum int
	// 店铺编号
	StoreId int64
	// 排序编号 --- 小的靠前
	Sort int
	// 商品规格 --- 如[{"name":"小白菜1斤","price":"2.00"},{"name":"小白菜2斤","price":"3.60"}]
	Format string
	//商品分类 --- 0商品，1服务
	Type bool
	//创建时间 --- 时间戳
	AddTime int64
}

//分页获取数据
func GetProductPage(index, size int, where string) ([]*TProduct, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TProduct{}
	dataCounts := []*TProduct{}
	//返回数据列表
	sql := fmt.Sprintf("select * from t_product %v  order by id desc limit %v offset %v", where, size, size*(index-1))
	beego.Info(sql)
	_, err := ormHelper.Raw(sql).QueryRows(&data)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	//返回计数
	sqlCount := fmt.Sprintf("select * from t_product  %v ", where)
	count, err1 := ormHelper.Raw(sqlCount).QueryRows(&dataCounts)
	if err1 != nil {
		fmt.Printf("error is %v\n", err1)
	}
	beego.Info(count)
	return data, int(count), err
}

func AddProduct(product *TProduct) (int64, error) {
	o := orm.NewOrm()
	product.AddTime = time.Now().Unix()
	productId, err := o.Insert(product)
	return productId, err
}
