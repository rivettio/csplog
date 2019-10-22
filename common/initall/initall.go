package initall

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/orm"
	"go.etcd.io/etcd/clientv3"
	"gopkg.in/olivere/elastic.v2"
)



var (
	MysqlClient orm.Ormer
	EtcdClient  *clientv3.Client
	KafkaClient sarama.SyncProducer
	EsClient    *elastic.Client
	LogConfAll  ConfAll
)



