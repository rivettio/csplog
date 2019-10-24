package server

import (
	"fmt"
	"rivettio/github-csplog/csplog/common/initall"
	"sync"

	"github.com/Shopify/sarama"
	"go.etcd.io/etcd/clientv3"
	"gopkg.in/olivere/elastic.v2"
)

var (
	err           error
	EtcdClient    *clientv3.Client
	KafkaConsumer sarama.Consumer
	EsClient      *elastic.Client
	ConfigAll     initall.ConfAll
	MutexLock     sync.Mutex

	wg sync.WaitGroup
)

func init() {
	if err = InitServer(); err != nil {
		panic(fmt.Sprintf("init Server failed, err:%v", err))
	}
}

func InitServer() (err error) {

	if ConfigAll, err = initall.InitConf(); err != nil {
		return
	}

	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}

	if KafkaConsumer, err = initall.InitKafkaConsumer(); err != nil {
		return
	}

	if EsClient, err = initall.InitES(); err != nil {
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
	WatchLogConfigEtcd(ConfigAll.EtcdConf.ConfigKey)
	wg.Wait()
}
