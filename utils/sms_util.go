package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SMSRequest 是发送短信的请求结构体
type SMSRequest struct {
	PhoneNumbers  string `json:"PhoneNumbers"`
	SignName      string `json:"SignName"`
	TemplateCode  string `json:"TemplateCode"`
	TemplateParam string `json:"TemplateParam"`
}

// SendMessage 发送短信
func SendMessage(signName, templateCode, phoneNumbers, param string) {
	url := "https://dysmsapi.aliyuncs.com"
	accessKeyId := "yourAccessKeyId"
	accessKeySecret := "yourAccessKeySecret"

	request := SMSRequest{
		PhoneNumbers:  phoneNumbers,
		SignName:      signName,
		TemplateCode:  templateCode,
		TemplateParam: fmt.Sprintf(`{"code":"%s"}`, param),
	}

	// 将请求结构体转为JSON
	requestBody, err := json.Marshal(request)
	if err != nil {
		fmt.Println("请求序列化错误:", err)
		return
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("创建请求错误:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-acs-accesskey-id", accessKeyId)
	req.Header.Set("x-acs-accesskey-secret", accessKeySecret)

	client := &http.Client{Timeout: 10 * time.Second}

	// 发送HTTP请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求错误:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("短信发送成功")
	} else {
		fmt.Println("短信发送失败:", resp.Status)
	}
}
