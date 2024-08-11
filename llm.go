package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Data struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func Chat(content string) string {
	configP := ReadConfig()
	apiKey := configP.LLM.ApiKey
	prompt := configP.LLM.Prompt
	baseUrl := configP.LLM.BaseUrl
	model := configP.LLM.Model

	if len(apiKey) == 0 || len(prompt) == 0 {
		log.Fatalln("错误：请检查配置文件")
	}

	url := baseUrl + "/chat/completions"

	// 创建数据体
	data := Data{
		Model: model,
		Messages: []Message{
			{Role: "user", Content: fmt.Sprintf(prompt, content)},
		},
	}
	log.Println("Data: ", data)

	// 将数据体序列化为 JSON 字节
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("Error marshaling data:", err)
	}

	// 设置请求头
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", apiKey),
		"Content-Type":  "application/json",
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln("Error creating request:", err)
	}

	// 设置请求头到请求对象
	req.Header.Set("Authorization", header["Authorization"])
	req.Header.Set("Content-Type", header["Content-Type"])

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error making request:", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading response body:", err)
	}

	// 解析 JSON 响应
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// 打印结果
	reply := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	return reply
}
