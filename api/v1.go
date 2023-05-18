package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/xiaoWay/go-wxmoring/model"
	"io/ioutil"
	"net/http"
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
