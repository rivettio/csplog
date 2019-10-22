package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/models"
	"go.etcd.io/etcd/clientv3"
	"rivettio/github-csplog/csplog/common/initall"
)

var (
	DB         orm.Ormer
	EtcdClient *clientv3.Client
	ConfigAll  initall.ConfAll
)

func init() {

	orm.RegisterModel(new(LogConfig))

	models.Connect()

	orm.RunSyncdb("default", false, true)

	DB = orm.NewOrm()

	if err := initAll(); err != nil {
		panic(fmt.Sprintln("init database failed, err:%v", err))
	}
}

func initAll() (err error) {
	if ConfigAll, err = initall.InitConf(); err != nil {
		return
	}

	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}
	if err = initall.InitLogs(); err != nil {
		return
	}
	return
}

