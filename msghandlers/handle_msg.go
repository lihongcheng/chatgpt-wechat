package msghandlers

import (
	"chatgpt-wechat/chatgpt"
	"github.com/eatmoreapple/openwechat"
	"log"
	"os"
	"strings"
)

func Draw(msg *openwechat.Message, msgStr, atText string) bool {
	botPrefix := chatgpt.Config["bot_prefix"].(string)
	cdrawPrefix := chatgpt.Config["draw_prefix"].(string)
	drawPrefix := botPrefix + cdrawPrefix
	if botPrefix != "" {
		if strings.HasPrefix(msgStr, botPrefix) {
			msgStr = botPrefix + strings.TrimSpace(strings.TrimPrefix(msgStr, botPrefix))
		}
	}
	if strings.HasPrefix(msgStr, drawPrefix) {
		drawText := strings.TrimPrefix(msgStr, drawPrefix)
		if "" != drawText {
			imgUrl := chatgpt.GetAiImg(drawText)
			if "" != imgUrl {
				img, _ := os.Open(imgUrl)
				defer img.Close()
				if atText != "" {
					_, _ = msg.ReplyText(atText)
				}
				_, err := msg.ReplyImage(img)
				if err != nil {
					log.Printf("ImagesGenerations response user error: %v \n", err)
				}
			}
			return true
		}
	}
	return false
}
