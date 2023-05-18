package main

import (
	"fmt"
	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/xiaoWay/go-wxmoring/config"
	"github.com/xiaoWay/go-wxmoring/model"
	"log"
)

/*
	发送数据的格式；
	{
		"appToken":"AT_xxx",
		"content":"Wxpusher祝你中秋节快乐!",
		"summary":"消息摘要",//消息摘要，显示在微信聊天页面或者模版消息卡片上，限制长度100，可以不传，不传默认截取content前面的内容。
		"contentType":1,//内容类型 1表示文字  2表示html(只发送body标签内部的数据即可，不包括body标签) 3表示markdown
		"topicIds":[ //发送目标的topicId，是一个数组！！！，也就是群发，使用uids单发的时候， 可以不传。
		123
		],
		"uids":[//发送目标的UID，是一个数组。注意uids和topicIds可以同时填写，也可以只填写一个。
		"UID_xxxx"
		],
		"url":"https://wxpusher.zjiecode.com", //原文链接，可选参数
		"verifyPay":false //是否验证订阅时间，true表示只推送给付费订阅用户，false表示推送的时候，不验证付费，不验证用户订阅到期时间，用户订阅过期了，也能收到。
	}
*/

/*
	{
		"code": 1000, //状态码
		"msg": "处理成功",//提示消息
		"data": [ //每个uid/topicid的发送状态，和发送的时候，一一对应，是一个数组，可能有多个
			{
				"uid": "UID_xxx",//用户uid
				"topicId": null, //主题ID
				"messageId": 121,//废弃⚠️，请不要再使用，后续会删除这个字段
				"messageContentId": 2123,//消息内容id，调用一次接口，生成一个，你可以通过此id调用删除消息接口，删除消息。本次发送的所有用户共享此消息内容。
				"sendRecordId": 12313,//消息发送id，每个uid用户或者topicId生成一个，可以通过这个id查询对某个用户的发送状态
				"code": 1000, //1000表示发送成功
				"status": "创建发送任务成功"
			}
		],
		"success": true
	}
*/

func sendTemplateMessage(msg string) {
	// msg: --> Token + content + contentType
	// todo: 这个官方的解析有点问题 func (m *Message) AddUId(id string, more ...string) *Message  优化一下；
	sender := model.NewMessage(config.AppConfig.WxPusher.AppToken).SetContent(msg).SetContentType(config.AppConfig.WxPusher.ContentType).SetSummary(config.AppConfig.WxPusher.Summary)
	sender.AddUId()
	msgArr, err := wxpusher.SendMessage(sender)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("To %s : %s", msgArr, msg)
}
