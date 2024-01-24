package msghandlers

import (
	"chatgpt-wechat/chatgpt"
	"github.com/eatmoreapple/openwechat"
)

var Session *chatgpt.Session

func init() {
	Session = chatgpt.NewSession()
}

func Handle(msg *openwechat.Message) {
	//私聊
	if msg.IsSendByFriend() {
		HandleFriend(msg)
	}
	//群聊
	if msg.IsSendByGroup() {
		HandleGroup(msg)
	}
}
