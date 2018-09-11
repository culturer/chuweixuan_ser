package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//商铺表
type TStore struct {
	//商铺编号
	Id int64
	//官方店铺
	IsOffice bool
	// 商铺名
	Name string
	//店铺图标
	Icon string
	//联系电话
	Tel string
	// 商铺简介
	Desc string
	// 商铺详情
	Content string `orm:"type(text);null"`
	// 商铺标签 --- 最多三个
	Tab string
	// 商铺地址
	Address string
	//商铺坐标(供定位使用)
	Position string
	//商铺所有者
	UserId int64
	//是否停业
	IsClose bool
	//停业告示
	CloseMsg string
	//上班时间
	OpenTime string
	//是否冻结
	IsLock bool
	//创建时间 --- 时间戳
	AddTime int64
	//商铺分类
	StoreTypeId int64
	// 小区编号
	CommunityId int64
}

//分页获取数据
func GetStorePage(index, size int, where string) ([]*TStore, int, error) {
	//orm模块
	ormHelper := orm.NewOrm()
	//返回data数据
	data := []*TStore{}
	dataCounts := []*TStore{}
	//返回数据列表
	sql := fmt.Sprintf("select * from t_store %v  order by id desc limit %v offset %v", where, size, size*(index-1))
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

//新建
func AddStore(store *TStore) (int64, error) {
	o := orm.NewOrm()
	storeId, err := o.Insert(store)
	return storeId, err
}

//新建用户
func UpdateStore(store *TStore) (int64, error) {
	o := orm.NewOrm()
	storeId, err := o.Update(store)
	return storeId, err
}

//查询
func GetStoreById(userId int64) (*TStore, error) {
	o := orm.NewOrm()
	store := new(TStore)
	qs := o.QueryTable("t_store")
	err := qs.Filter("user_id", userId).One(store)
	return store, err
}

//查询
func GetStoreByStoreId(userId int64) (*TStore, error) {
	o := orm.NewOrm()
	store := new(TStore)
	qs := o.QueryTable("t_store")
	err := qs.Filter("user_id", userId).One(store)
	return store, err
}

//查询
func GetOfficeByCommitId(communityId int64) (*TStore, error) {
	o := orm.NewOrm()
	store := new(TStore)
	qs := o.QueryTable("t_store")
	err := qs.Filter("community_id", communityId).Filter("is_office", true).One(store)
	return store, err
}
