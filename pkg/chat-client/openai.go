package chatclient

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"k8s.io/utils/env"
)

type OpenAIChatClient struct {
	client *openai.Client
	model  string
}

func NewOpenAIChatClient(model string) *OpenAIChatClient {
	apiKey := env.GetString("OPENAI_API_KEY", "")
	baseUrl := env.GetString("OPENAI_BASE_URL", "https://api.openai.com/v1")
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseUrl
	client := openai.NewClientWithConfig(config)
	return &OpenAIChatClient{client: client, model: model}
}

func (o *OpenAIChatClient) SendMessage(msg string) (string, error) {
	completion, err := o.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: o.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemMessage},
			{Role: openai.ChatMessageRoleUser, Content: msg},
		},
	})

	if err != nil {
		return "", err
	}

	return completion.Choices[0].Message.Content, nil
}
