package msghandlers

import (
	"chatgpt-wechat/chatgpt"
	"github.com/eatmoreapple/openwechat"
	"log"
	"strings"
)

func HandleFriend(msg *openwechat.Message) {
	sender, _ := msg.Sender()
	log.Printf("收到朋友[%s]的消息：%s\n", sender.NickName, msg.Content)
	//处理文本消息
	if msg.IsText() {
		//画图
		msg.Content = strings.TrimSpace(msg.Content)
		if Draw(msg, msg.Content, "") {
			return
		}
		botPrefix := chatgpt.Config["bot_prefix"].(string)
		if botPrefix == "" {
			//nothing to do
		} else {
			if strings.HasPrefix(msg.Content, botPrefix) {
				msg.Content = strings.TrimPrefix(msg.Content, botPrefix)
			} else {
				return
			}
		}
		aiMSg, err := chatgpt.GetAiMsg(msg, Session)
		if err != nil {
			_, _ = msg.ReplyText("我好像迷失在了无尽的宇宙中，你能帮我找到回家的路吗？")
			return
		}
		_, err = msg.ReplyText(aiMSg)
		if err != nil {
			return
		}
	} else if msg.IsVoice() {
		if chatgpt.Config["bot_prefix"].(string) != "" {
			return
		}
		text, err := chatgpt.VoiceToText(msg)
		if err != nil {
			log.Println("语音转文字失败:", err)
			_, _ = msg.ReplyText("我好像迷失在了无尽的宇宙中，你能帮我找到回家的路吗？")
			return
		}
		//画图
		if Draw(msg, text, "") {
			return
		}
		msg.Content = text
		aiMSg, err := chatgpt.GetAiMsg(msg, Session)
		if err != nil {
			_, _ = msg.ReplyText("我好像迷失在了无尽的宇宙中，你能帮我找到回家的路吗？")
			return
		}
		_, err = msg.ReplyText(aiMSg)
		if err != nil {
			return
		}
	}
}
