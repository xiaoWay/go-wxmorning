package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Mod            string         `json:"mod"`
	Cron           string         `json:"cron"`
	LoveDay        string         `json:"love_day"`
	BirthDay       string         `json:"birth_day"`
	WechatOfficial WechatOfficial `json:"wechat_official"`
	Baidutianqi    Baidutianqi    `json:"baidutianqi"`
	CaiHongPiKey   string         `json:"caihongpikey"`
}

// WechatOfficial 微信公众号配置
type WechatOfficial struct {
	AppID      string   `json:"app_id"`      // appid
	AppSecret  string   `json:"app_secret"`  // appsecret
	OpenIds    []string `json:"open_ids"`    // 要接受消息的人
	TemplateID string   `json:"template_id"` // 必须, 模版ID
}

// baidutianqi 百度天气配置
type Baidutianqi struct {
	Ak         string `json:"ak"`
	DistrictId string `json:"district_id"`
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
