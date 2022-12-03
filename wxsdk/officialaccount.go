package wxsdk

import (
	"github.com/silenceper/wechat/v2/officialaccount/broadcast"
	"github.com/silenceper/wechat/v2/officialaccount/datacube"
	"github.com/silenceper/wechat/v2/officialaccount/device"
	"github.com/silenceper/wechat/v2/officialaccount/draft"
	"github.com/silenceper/wechat/v2/officialaccount/freepublish"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/silenceper/wechat/v2/officialaccount/menu"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/ocr"
)

type OfficialAccount struct {
	ctx          *context.Context
	basic        *basic.Basic
	menu         *menu.Menu
	oauth        *oauth.Oauth
	material     *material.Material
	draft        *draft.Draft
	freepublish  *freepublish.FreePublish
	js           *js.Js
	user         *user.User
	templateMsg  *message.Template
	managerMsg   *message.Manager
	device       *device.Device
	broadcast    *broadcast.Broadcast
	datacube     *datacube.DataCube
	ocr          *ocr.OCR
	subscribeMsg *message.Subscribe
}
