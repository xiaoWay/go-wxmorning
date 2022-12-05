package main

//使用了
import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	officialConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/tidwall/gjson"
	"gomoring/config"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var official *officialaccount.OfficialAccount

func init() {
	//设置时区
	var cstZone = time.FixedZone("CST", 8*3600) //UTC-8
	time.Local = cstZone

	//config.Getconfig()
	//sdk 配置api通信 cache保存
	official = wechat.NewWechat().GetOfficialAccount(&officialConfig.Config{
		AppID:     config.AppConfig.WechatOfficial.AppID,
		AppSecret: config.AppConfig.WechatOfficial.AppSecret,
		Cache:     cache.NewMemory(),
	})

	//如果为test模式发送一条 就推迟os
	if config.AppConfig.Mod == "test" {
		log.Println("当前是测试模式，将立即发送一条消息并退出，如需定时发送请将 mod 值改为其他任意值，只要不是 test 就行")
		sendTemplateMessage()
		os.Exit(0)
	}
}

var cronEntryId cron.EntryID

func sendTemplateMessage() {
	//模板消息设置
	// data: {{riqi.DATA}} //2022-11-21 星期一
	// beizhu {{beizhu.DATA}}
	// weather 天气：{{tianqi.DATA}}} //晴
	// low 最低温度：{{low.DATA}} 度 //17
	// high 最高温度：{{high.DATA}} 度 //25
	// 今天是我们恋爱的第 {{lianai.DATA}}天 // 250
	// 距离你的生日还有{{shengri.DATA}}天// 251

	// {{caihongpi.DATA}}彩虹屁
	// {{jinju.DATA}} 金句
	text, low, high := getWeather() //天气，最低温度，最高温度
	now := time.Now()
	day := now.Format("2006-01-02")
	weekday := time.Now().Weekday()
	riqi := fmt.Sprintf("%s %s", day, weekday) // yyyy-mm-dd weekday
	for _, openId := range config.AppConfig.WechatOfficial.OpenIds {
		msg := &message.TemplateMessage{
			ToUser:     openId,
			TemplateID: config.AppConfig.WechatOfficial.TemplateID,
			Data: map[string]*message.TemplateDataItem{
				"riqi":      {Value: riqi, Color: randomcolor()},
				"tianqi":    {Value: text, Color: randomcolor()},
				"low":       {Value: low, Color: randomcolor()},
				"high":      {Value: high, Color: randomcolor()},
				"lianai":    {Value: fmt.Sprintf("%d", getLoverDay()), Color: randomcolor()},
				"shengri":   {Value: fmt.Sprintf("%d", getBirthDay()), Color: randomcolor()},
				"caihongpi": {Value: fmt.Sprintf("%s\n", getCaiHongPi()), Color: randomcolor()},
				"jinju":     {Value: getJinju(), Color: randomcolor()},
			},
		}
		_, err := official.GetTemplate().Send(msg)

		if err != nil {
			log.Printf("发送模版消息失败 openId=[%s] err:%s\n", openId, err.Error())
		}
		fmt.Println(msg) //后台输出函数，打印面板
		//fmt.Println("222")
		//懒 现在输出的是msg的地址 如果想要好的输出那就解析一下map就好了，然后对每个地址求值输出
		//我直接一股脑输出
	}
}

func main() {
	c := cron.New()
	var err error
	cronEntryId, err = c.AddFunc(config.AppConfig.Cron, func() {
		sendTemplateMessage()
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
	colorArr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"} //基础16进制串
	rand.Seed(time.Now().UnixNano())
	c := ""
	for i := 0; i < 6; i++ {
		c = c + colorArr[rand.Intn(16)]
	}
	return "#" + c

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
