package msghandlers

import (
	"chatgpt-wechat/chatgpt"
	"github.com/eatmoreapple/openwechat"
	"log"
	"strings"
)

func HandleGroup(msg *openwechat.Message) {
	if !msg.IsAt() {
		return
	}
	groupSender, _ := msg.SenderInGroup()
	atText := "@" + groupSender.NickName + " "
	sender, _ := msg.Sender()
	group := openwechat.Group{User: sender}
	log.Printf("收到群[%s]消息：%s\n", group.NickName, msg.Content)
	if msg.IsText() {
		atMeText := "@" + sender.Self().NickName
		requestText := strings.TrimSpace(strings.ReplaceAll(msg.Content, atMeText, ""))
		//画图
		if Draw(msg, requestText, atText) {
			return
		}
		botPrefix := chatgpt.Config["bot_prefix"].(string)
		if botPrefix == "" {
			//nothing to do
		} else {
			if strings.HasPrefix(requestText, botPrefix) {
				requestText = strings.TrimPrefix(requestText, botPrefix)
			} else {
				return
			}
		}
		msg.Content = requestText
		aiMSg, err := chatgpt.GetAiMsg(msg, Session)
		if err != nil {
			_, _ = msg.ReplyText(atText + "我好像迷失在了无尽的宇宙中，你能帮我找到回家的路吗？")
			return
		}
		_, err = msg.ReplyText(atText + aiMSg)
	}
}
