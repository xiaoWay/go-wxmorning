package service

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/xiaoWay/go-wxmoring/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// caiHongPi 使用彩虹屁api接口
func getCaiHongPi() string {
	url := fmt.Sprintf("http://api.tianapi.com/caihongpi/index?key=%s", config.AppConfig.CaiHongPiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	return gjson.GetBytes(data, "newslist.0.content").String()
}

// baidutianqi 使用百度天气api
func getWeather() (text string, low string, high string) {
	url := fmt.Sprintf("https://api.map.baidu.com/weather/v1/?district_id=%s&data_type=all&ak=%s", config.AppConfig.Baidutianqi.DistrictId, config.AppConfig.Baidutianqi.Ak)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("获取天气消息失败 err:", err.Error())
		os.Exit(0)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("获取天气消息失败 status:", resp.Status)
		os.Exit(0)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	text = gjson.GetBytes(data, "result.now.text").String()
	low = gjson.GetBytes(data, "result.forecasts.0.low").String()
	high = gjson.GetBytes(data, "result.forecasts.0.high").String()
	return text, low, high
}

// 金山词霸每日一句api 使用
func getJinju() string {
	url := fmt.Sprintf("http://open.iciba.com/dsapi/")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	content := gjson.GetBytes(data, "content").String()
	note := gjson.GetBytes(data, "note").String()
	return fmt.Sprintf("%s \n %s", content, note)
}

func getUrl() string {
	key := "zzandxw/1.jpg"
	url := GetDownload(key)
	return url
}

//// 即将废弃 Bing图片
//func getUrl() string {
//	resp, err := http.Get("https://api.vvhan.com/api/bing?type=json")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	defer resp.Body.Close()
//	data, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	url := gjson.GetBytes(data, "data.url").String()
//	return url
//}

// 获取恋爱多少天
func getLoverDay() int {
	c, err := time.Parse("2006-01-02", config.AppConfig.LoveDay)
	if err != nil {
		log.Panicf("parse error: %v", err)
	}
	return int(time.Now().Sub(c).Hours() / 24.0)
}

// 获取当前时间
func getCurrentDate() time.Time {
	nowStr := time.Now().Format("2006-01-02")
	now, _ := time.Parse("2006-01-02", nowStr)
	return now
}

// 判断还有多少天生日
func getBirthDay() int {
	birth, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%s", getCurrentDate().Year(), config.AppConfig.BirthDay))
	if err != nil {
		log.Println("GetBirthDay 错误", birth)
		return -1111
	}
	if getCurrentDate().Sub(birth) > 0 {
		birth, _ = time.Parse("2006-01-02", fmt.Sprintf("%d-%s", getCurrentDate().Year()+1, config.AppConfig.BirthDay))
	}
	return int(birth.Sub(getCurrentDate()).Hours() / 24.0)
}
