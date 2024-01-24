package chatgpt

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

type Session struct {
	cache *cache.Cache
}

func NewSession() *Session {
	return &Session{cache: cache.New(time.Minute*10, time.Minute*10)}
}

func (s *Session) GetCache(key string) []map[string]string {
	sessionContext, ok := s.cache.Get(key)
	if !ok {
		return []map[string]string{}
	}
	// 定义一个切片，用于存储解析后的数据
	var qa []map[string]string
	// 将 JSON 字符串解析为切片
	err := json.Unmarshal([]byte(sessionContext.(string)), &qa)
	if err != nil {
		log.Println("GetCache Error:", err)
		return []map[string]string{}
	}
	return qa
}

func (s *Session) SetCache(key string, question, reply string) error {
	qa := []map[string]string{
		{
			"role":    "user",
			"content": question,
		}, {
			"role":    "assistant",
			"content": reply,
		},
	}
	result := []map[string]string{}
	historyMsg := s.GetCache(key)
	//由于token限制,只保留10轮对话
	rounds := Config["session_rounds"].(int)
	limit := rounds*2 - 2
	if len(historyMsg) >= limit {
		result = historyMsg[len(historyMsg)-limit:]
		result = append(result, qa...)
	} else {
		if len(historyMsg) > 0 {
			historyMsg = append(historyMsg, qa...)
			result = append(result, historyMsg...)
		} else {
			result = append(result, qa...)
		}
	}
	// 将切片转为 JSON 字符串
	jsonString, err := json.Marshal(result)
	if err != nil {
		log.Println("SetCache Error:", err)
		return err
	}
	s.cache.Set(key, string(jsonString), cache.DefaultExpiration)
	return nil
}
