# 项目名称

**微信聊天接入ChatGPT**

## 简介

该项目实现了将微信接入 ChatGPT，使得编译后的程序可以在本地和云服务器上运行，而无需科学上网。该聊天机器人支持微信私聊、群聊，以及提供画图和语音回复功能。

## 功能特点

- **微信接入：** 实现微信私聊和群聊的机器人功能。
- **本地和云服务器运行：** 支持在本地和云服务器上编译和运行。
- **无需科学上网：** 能够在没有科学上网的环境中正常运行。
- **画图功能：** 机器人支持画图功能，使得用户可以通过微信发送指令绘制简单图形。
- **语音回复：** 提供语音回复功能，使得机器人可以通过语音与用户进行交互。

## 快速开始

### 1. 克隆仓库
```bash
git clone https://github.com/lihongcheng/chatgpt-wechat.git
cd chatgpt-wechat
```
### 2. 编译程序
```bash
go build -o chatgpt-wechat
```
### 3. 运行程序
```bash
./chatgpt-wechat your-openai-key（必传） 角色设定（可选参数）
```

## 配置文件
```bash
package chatgpt

var Config = map[string]interface{}{
	"chat_api":       "https://chatgpt.codeworld.top/v1/chat/completions",//这里是ChatGPT的API代理地址
	"voice_api":      "https://chatgpt.codeworld.top/v1/audio/transcriptions",//这里是语音转文本的API代理地址
	"draw_api":       "https://chatgpt.codeworld.top/v1/images/generations",//这里是画图的API代理地址
	"session_rounds": 10,//这里是上下文最大对话轮数
	"role":           "你是ChatGPT, 一个由OpenAI训练的大型语言模型, 你旨在回答并解决人们的任何问题，并且可以使用多种语言与人交流。",//角色设定,可传给命令行第二个参数
	"draw_prefix":    "画",//画图命令前缀
	"bot_prefix":     "",//机器人命令前缀
}
```
## 使用示例

### 1.启动程序
```bash
./chatgpt-wechat your-openai-key（必传参数） 角色设定（可选参数）
```
### 2.扫描微信二维码登录

### 3. 与机器人聊天

- 支持微信群聊@回复
  
![image](https://github.com/lihongcheng/aichat/assets/20829680/f6acd473-4a6f-4171-8688-4bf8039a3d9d)

- 支持微信私聊

 ![image](https://github.com/lihongcheng/aichat/assets/20829680/fc874ce5-eedc-4096-9ff8-56e72e75ff0a)

- 支持AI绘图
  
  ![image](https://github.com/lihongcheng/chatgpt-wechat/assets/20829680/9fe8816f-2465-4791-a1a7-76cb769c31d2)

- 支持微信语音自动识别

  ![image](https://github.com/lihongcheng/aichat/assets/20829680/cbf0fe15-3441-4eea-8dfe-55df6632247b)

- 支持会话上下文记忆功能
  
  ![image](https://github.com/lihongcheng/chatgpt-wechat/assets/20829680/efd70e94-86b9-4c05-91b4-f933eb5a07ec)


