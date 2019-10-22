package initall

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
)

func InitEtcd() (etcdClient *clientv3.Client, err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   LogConfAll.EtcdConf.EtcdAddr,
		DialTimeout: time.Duration(LogConfAll.EtcdConf.DailTimeout) * time.Second,
	})
	if err != nil {
		err = fmt.Errorf("connect etcd failed, err:", err)
		logs.Error(err)
		return
	}

	etcdClient = cli
	logs.Debug("init etcd success")
	return
}
