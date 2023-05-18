package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xiaoWay/go-wxmoring/model"
	"io/ioutil"
	"net/http"
	"time"
)

// URLBase 接口域名
const URLBase = "http://wxpusher.zjiecode.com"

// URLSendMessage 发送消息
const URLSendMessage = URLBase + "/api/send/message"

// SendMessage 发送消息
func SendMessage(message *model.Message) ([]model.SendMsgResult, error) {
	msgResults := make([]model.SendMsgResult, 0)
	// 校验消息内容
	err := message.Check()
	if err != nil {
		return msgResults, err
	}
	// 参数转json
	requestBody, _ := json.Marshal(message)
	resp, err := http.Post(URLSendMessage, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	defer func() { _ = resp.Body.Close() }()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	result := model.Result{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	if !result.Success() {
		return msgResults, model.NewError(result.Code, errors.New(result.Msg))
	}
	// result.Data 转为[]model.SendMsgResult
	byteData, err := json.Marshal(result.Data)
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	err = json.Unmarshal(byteData, &msgResults)
	if err != nil {
		return msgResults, model.NewSDKError(err)
	}
	return msgResults, nil
}

func InitMsg() string {
	text, low, high := getWeather()

	riqi := time.Now().Format("2006-01-02 Monday")
	//now := time.Now()
	//day := now.Format("2006-01-02")
	//weekday := time.Now().Weekday()
	//riqi := fmt.Sprintf("%s %s", day, weekday) // yyyy-mm-dd weekday
	// 格式化生成的文字为Markdown形式
	output := fmt.Sprintf(
		"<span style='color:%s'>%s</span>\n\n"+
			"天气：<span style='color:%s'> %s </span>\n"+
			"最低温度：<span style='color:%s'> %s </span>\n"+
			"最高温度：<span style='color:%s'> %s </span>\n"+
			"今天是我们恋爱的第: <span style='color:%s'>%d</span> 天\n"+
			"距离你的生日还有: <span style='color:%s'>%d</span>天\n\n"+
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
