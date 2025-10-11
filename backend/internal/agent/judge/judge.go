// Package judge judge.go
//
// 功能:
// - 提交题目和用户作答交由Ai判题
//
// 作者: LuckyQu
// 创建日期: 2025-10-11
// 修改日期: 2025-10-11
package judge

import (
	"backend/configs"
	"backend/internal/agent"
	agentModel "backend/internal/agent/model"
	"backend/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// SendToJudge 交由Ai进行判题
func SendToJudge(question *model.Question, answer string) (*agentModel.JudgeResult, error) {
	// 获取题目参数
	questionTitle := question.Title
	questionContent := question.Content
	// 构建请求体
	requestBody := &agentModel.Request{
		Model: configs.Config.Agent.Model,
		Input: agentModel.Input{
			Messages: []agentModel.Message{
				{
					// 初始化判题身份
					Role:    "system",
					Content: "你现在作为一名判题官，请判题后严格返回{result:true|false, suggestion:正确则补充(如有)，错误则指出改正}的JSON结构",
				},
				{
					// 拼接消息
					Role:    "user",
					Content: "问题的题目是:\n" + questionTitle + "问题的题干是\n" + questionContent + "用户给出的答案是:\n" + answer,
				},
			},
		},
		Parameters: agentModel.Parameters{
			EnableThinking: configs.Config.Agent.EnableThinking,
			MaxTokens:      configs.Config.Agent.MaxTokens,
			ResponseFormat: agentModel.ResponseFormat{
				Type: configs.Config.Agent.ResponseFormatType,
			},
		},
		MaxInputTokens: configs.Config.Agent.MaxInputTokens,
	}
	// 序列化为JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	// 构建请求
	request, err := http.NewRequest("POST", configs.Config.Agent.BaseUrl, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, err
	}
	// 编辑请求头
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+configs.Config.Agent.ApiKey)
	// 发送请求
	response, err := agent.GetAgentClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	// 拿到响应
	responseBody := agentModel.Response{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}
	judgeResult := agentModel.JudgeResult{}
	// 反序列化得到结果
	err = json.Unmarshal([]byte(responseBody.Output.Text), &judgeResult)
	if err != nil {
		return nil, err
	}
	return &judgeResult, nil
}
