package chatgpt

var Config = map[string]interface{}{
	"chat_api":       "https://chatgpt.codeworld.top/v1/chat/completions",              //这里是ChatGPT的API代理地址
	"voice_api":      "https://chatgpt.codeworld.top/v1/audio/transcriptions",          //这里是语音转文本的API代理地址
	"draw_api":       "https://chatgpt.codeworld.top/v1/images/generations",            //这里是画图的API代理地址
	"session_rounds": 10,                                                               //这里是上下文最大对话轮数
	"role":           "你是ChatGPT, 一个由OpenAI训练的大型语言模型, 你旨在回答并解决人们的任何问题，并且可以使用多种语言与人交流。", //角色设定,可传给命令行第二个参数
	"draw_prefix":    "画",                                                              //画图命令前缀
	"bot_prefix":     "",                                                               //机器人命令前缀
	"model":          "gpt-3.5-turbo-16k",                                              //这里是ChatGPT的模型版本
}
