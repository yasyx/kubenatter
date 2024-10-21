package chatclient

import "github.com/sashabaranov/go-openai"

var systemMessage = `
You are a k8s administrator, you can help me to manage the k8s resources.
When you generate a k8s resource,Do not output any content except for YAML content, and do not place YAML inside code blocks.
`

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
