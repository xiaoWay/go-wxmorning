package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	Mod            string         `json:"mod"`
	Cron           string         `json:"cron"`
	LoveStartDate  string         `json:"love_start_date"`
	LoveStartTime  time.Time      `json:"love_start_time"`
	BirthDate      string         `json:"birth_date"`
	BirthTime      time.Time      `json:"birth_time"`
	WechatOfficial WechatOfficial `json:"wechat_official"`
	Baidutianqi    Baidutianqi    `json:"baidutianqi"`
	CaiHongPiKey   string         `json:"caihongpikey"`

	Colors Colors `json:"colors"`
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

type Colors struct {
	Qinghua string `json:"qinghua"`
}

var DefaultConfig = &Config{
	Colors: Colors{
		Qinghua: "#550038",
	},
}

// 获取恋爱多少天
func (c *Config) GetLoverDay() int {
	return int(time.Now().Sub(c.LoveStartTime).Hours() / 24.0)
}

func getCurrentDate() time.Time {
	nowStr := time.Now().Format("2006-01-02")
	now, _ := time.Parse("2006-01-02", nowStr)
	return now
}

// 判断还有多少天生日
func (c *Config) GetBirthDay() int {
	birth, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%s", getCurrentDate().Year(), c.BirthDate))
	if err != nil {
		log.Println("GetBirthDay 错误", birth)
		return -1111
	}
	if getCurrentDate().Sub(birth) > 0 {
		birth, _ = time.Parse("2006-01-02", fmt.Sprintf("%d-%s", getCurrentDate().Year()+1, c.BirthDate))
	}
	return int(birth.Sub(getCurrentDate()).Hours() / 24.0)
}

// 配置自检
func Parse() {
	configBytes, err := ioutil.ReadFile("./config_local.json")
	if err != nil {
		configBytes, err = ioutil.ReadFile("./config.json")
		if err != nil {
			log.Panicf("解析配置文件出错 err: %s", err.Error())
		}
	}
	if err := json.Unmarshal(configBytes, DefaultConfig); err != nil {
		log.Panicf("解析配置文件 Unmarshal 出错  err: %s", err.Error())
	}
	if DefaultConfig.LoveStartTime, err = time.Parse("2006-01-02", DefaultConfig.LoveStartDate); err != nil {
		log.Panicf("解析配置文件 love_start_date 出错  err: %s", err.Error())
	}
	if DefaultConfig.BirthTime, err = time.Parse("01-02", DefaultConfig.BirthDate); err != nil {
		log.Panicf("解析配置文件 birth_date 出错  err: %s", err.Error())
	}
}
