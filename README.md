# 微信测试号——给女朋友的每日早安提醒
每天早上给女朋友发个早安提醒。男码友专属示爱神器

欢迎issue和clone

本程序基于golang开发，使用了[silenceper/wechat/v2](https://github.com/silenceper/wechat) go sdk工具包，百度天气Api，彩虹屁Api，每日一言Api

效果如下：

![image-20221203184324368](readmefiles/1670064199037.png)

### 第一步

到[微信测试号](https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login)接口申请一个账号，扫码登录就可以获取，可以得到appID，appsecret的信息



呼叫你的女朋友🚺扫个码，能在后台获取到她的openid。

程序需要到[彩虹屁](https://www.tianapi.com/apiview/181)申请一个自己的api，获得一个自己的key

天气接口使用的是[百度天气](https://lbsyun.baidu.com/index.php?title=webapi/weather)，需要获取api的ak和女朋友所在地区的district_id：一般为当地的身份证前六位

------



### 第二步

填写config.json 

```
{
  "mod": "",  //test为测试模式发送一条消息即关闭，prod为生产模式
  "cron": "30 7 * * *", //定时开启任务 每天的7：30
  "love_start_date": "xxxx-xx-xx",  //恋爱开始的日子 yyyy-mm-dd
  "birth_date": "xx-xx",  //生日的月 日  yy-dd
  "wechat_official": {
    "app_id": "",   //自己的appid
    "app_secret": "",  //自己的secretid
    "open_ids": [
      ""  她的openid
    ],
    "template_id": ""  //创建消息模板返回的模板id
  },
  "baidutianqi": {
    "ak": "", "" //百度天气ak
    "district_id": ""  //所在地区district_id
  },
  "colors": {
    "qinghua": "#550038" //默认颜色
  },
  "caihongpikey": "" //彩虹皮key
}
```

------



### 第三步

填写在微信测试号官网模板信息

![image-20221203203142390](readmefiles/1670070683032.png)

模板内容：

```
{{riqi.DATA}}  

{{beizhu.DATA}}
天气：{{tianqi.DATA}}
最低温度：{{low.DATA}} 度
最高温度：{{high.DATA}} 度
今天是我们恋爱的第 {{lianai.DATA}} 天
距离你的生日还有 {{shengri.DATA}} 天
{{caihongpi.DATA}} 

{{jinju.DATA}}
```

------



### 第四步

build然后运行，大功完成   **编译好的exe要和config.json要放在一起**

```go
    go build main.go
```

