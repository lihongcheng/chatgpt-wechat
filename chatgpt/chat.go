package chatgpt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func GetAiMsg(msg *openwechat.Message, session *Session) (string, error) {
	url := Config["chat_api"].(string)
	apiKey := os.Args[1]
	role := Config["role"].(string)
	if len(os.Args) >= 3 {
		role = os.Args[2]
	}

	messages := []map[string]string{
		{
			"role":    "system",
			"content": role,
		},
	}

	sender, _ := msg.Sender()
	historyMsg := session.GetCache(sender.AvatarID())
	if len(historyMsg) > 0 {
		messages = append(messages, historyMsg...)
	}
	messages = append(messages, map[string]string{
		"role":    "user",
		"content": msg.Content,
	})

	data := map[string]interface{}{
		"model":    Config["model"].(string),
		"messages": messages,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return "", err
	}
	//格式化打印请求体
	printJsonString, _ := json.MarshalIndent(data, "", "    ")
	log.Println("Request Body:", string(printJsonString))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return "", err
	}
	defer resp.Body.Close()

	log.Println("Response Status:", resp.Status)

	// 读取响应体
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println("Response Body:", buf.String())
	var responseBody map[string]interface{}
	err = json.Unmarshal(buf.Bytes(), &responseBody)
	if err != nil {
		return "", err
	}
	// 提取所需内容
	content := ""
	if choices, ok := responseBody["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := choice["message"].(map[string]interface{}); ok {
				if content, ok = message["content"].(string); ok {
					log.Println("Extracted Content:", content)
				}
			}
		}
	}
	if content != "" {
		err := session.SetCache(sender.AvatarID(), msg.Content, content)
		if err != nil {
			return "", err
		}
	}
	return content, nil
}

func GetAiImg(desc string) string {
	apiKey := os.Args[1]
	url := Config["draw_api"].(string)
	contentType := "application/json"
	data := map[string]interface{}{
		"prompt": desc,
		"n":      1,
		"size":   "512x512",
	}
	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("ImagesGenerations error1 %v", err)
		return ""
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ImagesGenerations error2 %v", err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("ImagesGenerations resp body %v", string(body))
	type ApiResponse struct {
		Created int `json:"created"`
		Data    []struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	var apiResp ApiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		log.Printf("ImagesGenerations error3 %v", err)
		return ""
	}
	if len(apiResp.Data) > 0 {
		imageURL := apiResp.Data[0].URL
		if _, err := os.Stat("images"); os.IsNotExist(err) {
			// 目录不存在，创建目录
			err := os.MkdirAll("images", 0755)
			if err != nil {
				log.Println("创建目录images失败:", err)
			}
			log.Println("images目录创建成功!")
		} else {
			// 目录已存在
			log.Println("images目录已存在")
		}
		// Generate random file name
		rand.Seed(time.Now().UnixNano())
		fileName := fmt.Sprintf("%d.png", rand.Int())
		fileName = "images/" + fileName
		// Download image
		imgResp, err := http.Get(imageURL)
		if err != nil {
			log.Printf("ImagesGenerations error4 %v", err)
			return ""
		}

		defer imgResp.Body.Close()

		// Save image to file
		file, err := os.Create(fileName)
		if err != nil {
			log.Printf("ImagesGenerations error5 %v", err)
			return ""
		}

		defer file.Close()

		imgBody, err := ioutil.ReadAll(imgResp.Body)
		if err != nil {
			log.Printf("ImagesGenerations error6 %v", err)
			return ""
		}
		_, err = file.Write(imgBody)
		if err != nil {
			log.Printf("ImagesGenerations error7 %v", err)
			return ""
		}
		log.Printf("Image saved as %v", fileName)
		return fileName
	} else {
		return ""
	}
}

func VoiceToText(msg *openwechat.Message) (string, error) {
	type VoiceResponse struct {
		Text string `json:"text"`
	}
	resp, _ := msg.GetVoice()
	// 检查响应头部 Content-Type 字段是否为 audio/mp3
	contentType := resp.Header.Get("Content-Type")
	if contentType == "audio/mp3" {
		// 生成随机字节数组
		randomBytes := make([]byte, 32)
		_, err := rand.Read(randomBytes)
		if err != nil {
			return "", err
		}
		// 将字节数组转换成 Base64 编码字符串
		randomString := base64.URLEncoding.EncodeToString(randomBytes)
		// 将响应正文保存到本地 MP3 文件
		// 判断目录是否存在
		if _, err := os.Stat("voice"); os.IsNotExist(err) {
			// 目录不存在，创建目录
			err := os.MkdirAll("voice", 0755)
			if err != nil {
				log.Println("创建目录voice失败:", err)
			}
			log.Println("voice目录创建成功!")
		} else {
			// 目录已存在
			log.Println("voice目录已存在")
		}
		audioFile := "voice/" + randomString + ".mp3"
		file, err := os.Create(audioFile)
		if err != nil {
			return "", err
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return "", err
		}
		log.Printf("MP3 file saved to %v", audioFile)
		apiKey := os.Args[1] // 替换成你的OpenAI API Key
		// 读取音频文件
		audioData, err := ioutil.ReadFile(audioFile)
		if err != nil {
			return "", err
		}
		// Create a new multipart/form-data request
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		// Add the audio file to the request as an UploadFile object
		part, err := writer.CreateFormFile("file", audioFile)
		if err != nil {
			return "", err
		}
		part.Write(audioData)
		// Add the model parameter to the request
		_ = writer.WriteField("model", "whisper-1")
		// Close the multipart writer to finalize the request body
		err = writer.Close()
		if err != nil {
			return "", err
		}
		request, err := http.NewRequest("POST", Config["voice_api"].(string), body)
		if err != nil {
			return "", err
		}
		request.Header.Set("Content-Type", writer.FormDataContentType())
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

		// 发送API请求并解析响应
		client := http.Client{}
		response, err := client.Do(request)

		// 打印响应头信息
		log.Println("Response Header:")
		for name, values := range response.Header {
			for _, value := range values {
				log.Printf("%s: %s\n", name, value)
			}
		}
		if err != nil {
			return "", err
		}
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		// 打印响应体信息
		log.Printf("Response Body:")
		log.Printf(string(responseBody))

		var transcription VoiceResponse
		err = json.Unmarshal(responseBody, &transcription)
		if err != nil {
			return "", err
		}
		// 输出识别结果
		log.Printf("voice to text:%v", transcription.Text)
		return transcription.Text, nil
	}
	return "", nil
}
