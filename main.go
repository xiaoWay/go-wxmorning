package main

//使用了
import (
	"github.com/robfig/cron/v3"
	"github.com/xiaoWay/go-wxmoring/api"
	"github.com/xiaoWay/go-wxmoring/config"
	"github.com/xiaoWay/go-wxmoring/service"
	"log"
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
		api.SendTemplateMessage(service.InitMsg())
		os.Exit(0)
	}
}

func main() {
	c := cron.New()
	var err error
	cronEntryId, err = c.AddFunc(config.AppConfig.Cron, func() {
		api.SendTemplateMessage(service.InitMsg())
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
