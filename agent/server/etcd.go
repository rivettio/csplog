package server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

type CollectionConf struct {
	Id       int
	LogPath  string
	Topic    string
	Status   int
	SendRate int
}

var (
	LogConfList []CollectionConf
	TailInfoMap map[int]TailInfo
)

func LoadLogConfigFromEtcd(etcdKey string) (list []CollectionConf, err error) {

	getTimeOut := time.Second * time.Duration(ConfigAll.EtcdConf.GetTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), getTimeOut)
	defer func(cancel context.CancelFunc) { cancel() }(cancel)

	logConfigInfo, err := EtcdClient.Get(ctx, etcdKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err : %v", etcdKey, err)
		return
	}

	for _, v := range logConfigInfo.Kvs {
		err = json.Unmarshal(v.Value, &list)
		if err != nil {
			logs.Error("Unmarshal logConfigInfo failed, err : %v", err)
			return
		}
	}

	TailInfoMap = SwitchoverTailInfoMap(list)

	for _, v := range TailInfoMap {

		CollectionStart(v)
	}
	return
}

func WatchLogConfigEtcd(EtcdKey string) {
	fmt.Println("watch etcd success")
	var err error
	for {

		var (
			logConfList  []CollectionConf
			watchSuccess = true
		)
		WatchChan := EtcdClient.Watch(context.Background(), EtcdKey)
		for WatchResponse := range WatchChan {
			for _, Event := range WatchResponse.Events {
				if Event.Type == 1 {
					logs.Warn("key[%s] 's config deleted", EtcdKey)
					continue
				}
				if Event.Type == 0 && string(Event.Kv.Key) == EtcdKey {
					err = json.Unmarshal(Event.Kv.Value, &logConfList)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						watchSuccess = false
						continue
					}
				}
			}

			if watchSuccess {
				logs.Debug("get config from etcd succ, %v", logConfList)
				MutexLock.Lock()
				LogConfList = logConfList
				MutexLock.Unlock()
				newMap := SwitchoverTailInfoMap(logConfList)
				oldMap := TailInfoMap

				for k, v := range newMap {
					if _, ok := oldMap[k]; ok {
						continue
					}
					CollectionStart(v)
				}

				for k, v := range oldMap {
					if _, ok := newMap[k]; !ok {
						v.ExitSign = true
					}
				}

				MutexLock.Lock()
				TailInfoMap = newMap
				MutexLock.Unlock()
			}
		}
	}
}

func SwitchoverTailInfoMap(logConfList []CollectionConf) (tailInfoMap map[int]TailInfo) {
	tailInfoMap = make(map[int]TailInfo)
	for _, v := range logConfList {
		tailInfoMap[v.Id] = TailInfo{
			LogConf: v,
		}
	}
	return
}

func CollectionStart(t TailInfo) {
	CollectionStartNode(t)
}

func CollectionStartNode(t TailInfo) {

	wg.Add(1)
	go TailStart(t)

	go t.SendToKafka()
}
