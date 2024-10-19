package chatclient

import "github.com/sashabaranov/go-openai"

type ChatClient interface {
	SendMessage(message string) (string, error)
}

func NewChatClient(model string) ChatClient {
	switch model {
	case openai.GPT4o, openai.GPT4oMini:
		return NewOpenAIChatClient(model)
	default:
		return nil
	}
}
