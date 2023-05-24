package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Mod          string      `json:"mod"`
	Cron         string      `json:"cron"`
	LoveDay      string      `json:"love_day"`
	BirthDay     string      `json:"birth_day"`
	WxPusher     WxPusher    `json:"wx_pusher"`
	Baidutianqi  Baidutianqi `json:"baidutianqi"`
	CaiHongPiKey string      `json:"caihongpikey"`
	Qiniu        Qiniu       `json:"qiniu"`
}

type WxPusher struct {
	AppToken    string   `json:"app_token"`    // apptoken
	Uids        []string `json:"uids"`         // 发送用户的Uid
	ContentType int      `json:"content_type"` // 1表示文字  2表示html(只发送body标签内部的数据即可，不包括body标签) 3表示markdown
	Summary     string   `json:"summary"`
}

// baidutianqi 百度天气配置
type Baidutianqi struct {
	Ak         string `json:"ak"`
	DistrictId string `json:"district_id"`
}

// qiniu
type Qiniu struct {
	AccessKey string `json:"accesskey"`
	SecretKey string `json:"secretkey"`
}

// 输出的默认配置
var AppConfig Config

// 配置自检
func init() {
	configBytes, err := ioutil.ReadFile("./config.json") //读jason 配置
	if err != nil {
		log.Panicf("解析配置文件出错 err: %s", err.Error()) //读不到返回配置错误信息
	}
	//反解config json
	if err := json.Unmarshal(configBytes, &AppConfig); err != nil {
		log.Panicf("解析配置文件 Unmarshal 出错  err: %s", err.Error())
	}
}

//// init前调用APPConfig 初始化
//func Getconfig() Config {
//	return AppConfig
//}
