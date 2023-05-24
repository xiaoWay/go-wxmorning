package service

import (
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/xiaoWay/go-wxmoring/config"
	"log"
	"math/rand"
	"time"
)

//授权凭证的签算需要七牛账号下的一对有效的Access Key和Secret Ke

func GetDownload() string {

	mac := qbox.NewMac(config.AppConfig.Qiniu.AccessKey, config.AppConfig.Qiniu.SecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	bucket := "xiaoway"
	limit := 1000
	prefix := "zzandxw/"
	delimiter := ""
	//初始列举marker为空
	marker := ""
	entries, _, _, _, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
	if err != nil {
		log.Fatal(err)
	}
	//for {
	//	entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
	//	if err != nil {
	//		fmt.Println("list error,", err)
	//		break
	//	}
	//	//print entries
	//	for _, entry := range entries {
	//		fmt.Println(entry.Key)
	//	}
	//	if hasNext {
	//		marker = nextMarker
	//	} else {
	//		//list end
	//		break
	//	}
	//}

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(entries))
	key := entries[randIndex].Key

	domain := "pic.xiaoway.cc"

	deadline := time.Now().Add(time.Second * 86400).Unix() //24h有效
	paivateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	url := fmt.Sprintf("http://"+"%s", paivateAccessURL)
	return url
}
