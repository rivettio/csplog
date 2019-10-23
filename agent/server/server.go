package server

import (
	"fmt"
	"github.com/Shopify/sarama"
	"go.etcd.io/etcd/clientv3"
	"rivettio/github-csplog/csplog/common/initall"
	"sync"
)

var (
	err         error
	EtcdClient  *clientv3.Client
	KafkaClient sarama.SyncProducer
	ConfigAll   initall.ConfAll
	MutexLock   sync.Mutex
	wg          sync.WaitGroup
)

func init() {
	if err = InitServer(); err != nil {
		panic(fmt.Sprintf("init Server failed, err:%v", err))
	}
}

func InitServer() (err error) {

	if ConfigAll, err = initall.InitConf(); err != nil {
		fmt.Println("InitConf error : ", err)
		return
	}

	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}

	if KafkaClient, err = initall.InitKafka(); err != nil {
		return
	}

	if err = initall.InitLogs(); err != nil {
		return
	}

	if LogConfList, err = LoadLogConfigFromEtcd(ConfigAll.EtcdConf.ConfigKey); err != nil {
		err = fmt.Errorf("LoadLogConfigFromEtcd err : %v", err)
		return
	}
	return
}

func Run() {

	go WatchLogConfigEtcd(ConfigAll.EtcdConf.ConfigKey)
	wg.Wait()
}
