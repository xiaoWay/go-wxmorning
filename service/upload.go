package service

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/xiaoWay/go-wxmoring/config"
	"time"
)

//授权凭证的签算需要七牛账号下的一对有效的Access Key和Secret Ke

func GetDownload(key string) string {
	mac := qbox.NewMac(config.AppConfig.QiNiu.AccessKey, config.AppConfig.QiNiu.SecretKey)
	domain := "pic.xiaoway.cc/"

	deadline := time.Now().Add(time.Second * 36000).Unix() //10h有效
	paivateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return paivateAccessURL
}
