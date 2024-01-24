package main

import (
	"chatgpt-wechat/msghandlers"
	"github.com/eatmoreapple/openwechat"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		log.Println("请传入openai的key")
		return
	}
	// 创建一个新的微信客户端实例
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
	// 注册消息处理函数
	bot.MessageHandler = msghandlers.Handle
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	//登录
	err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		log.Println(err)
		return
	}
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	err = bot.Block()
	if err != nil {
		log.Println(err)
		return
	}
}
