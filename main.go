package main

//使用了
import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/tidwall/gjson"
	"github.com/xiaoWay/go-wxmoring/config"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var cronEntryId cron.EntryID

func init() {
	//设置时区
	var cstZone = time.FixedZone("CST", 8*3600) //UTC-8
	time.Local = cstZone

	//如果为test模式发送一条 就推迟os
	if config.AppConfig.Mod == "test" {
		log.Println("当前是测试模式，将立即发送一条消息并退出，如需定时发送请将 mod 值改为其他任意值，只要不是 test 就行")
		sendTemplateMessage(createMsg())
		os.Exit(0)
	}
}

func main() {
	c := cron.New()
	var err error
	cronEntryId, err = c.AddFunc(config.AppConfig.Cron, func() {
		sendTemplateMessage(createMsg())
		log.Println("执行成功, 下次执行时间是", c.Entry(cronEntryId).Next.String())
	})
	if err != nil {
		log.Panicf("添加定时任务失败 err:%s", err.Error())
	}
	c.Start()
	log.Println("启动成功, 下次执行时间是", c.Entry(cronEntryId).Next.String())
	//阻塞
	select {}
}

func createMsg() string {

	// todo: 备注
	text, low, high := getWeather()
	now := time.Now()
	day := now.Format("2006-01-02")
	weekday := time.Now().Weekday()
	riqi := fmt.Sprintf("%s %s", day, weekday) // yyyy-mm-dd weekday
	// 格式化生成的文字为Markdown形式
	output := fmt.Sprintf(
		"<span style='color:%s'>%s</span>\n\n"+
			"天气：<span style='color:%s'> %s </span>\n"+
			"最低温度：<span style='color:%s'> %s </span>\n"+
			"最高温度：<span style='color:%s'> %s </span>\n"+
			"今天是我们恋爱的第 <span style='color:%s'>%d</span> 天\n"+
			"距离你的生日还有<span style='color:%s'>%d</span>天\n\n"+
			"<span style='color:%s'>%s</span>\n\n"+
			"<span style='color:%s'>%s</span>",
		randomcolor(), riqi,
		randomcolor(), text,
		randomcolor(), low,
		randomcolor(), high,
		randomcolor(), getLoverDay(),
		randomcolor(), getBirthDay(),
		randomcolor(), getCaiHongPi(),
		randomcolor(), getJinju(),
	)

	// 输出Markdown字符串
	return output
}

/*
方法区
func
func
方法区
*/
// caiHongPi 使用彩虹屁api接口

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

// 随机颜色函数设置
func randomcolor() string {
	rand.Seed(time.Now().UnixNano())
	R := uint8(rand.Intn(255))
	G := uint8(rand.Intn(255))
	B := uint8(rand.Intn(255))
	color := fmt.Sprintf("#%02x%02x%02x", R, G, B)
	return color
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
