package chatclient

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type FunctionCallParams struct {
	// FunctionCallParams is the parameters for the function call.
	// The structure of the parameters depends on the function definition.
	Parameters map[string]interface{}
}

type FunctionCall struct {
	FunctionDefinition openai.FunctionDefinition
	Func               func(params map[string]interface{}) (map[string]interface{}, error)
}

var f3 = openai.FunctionDefinition{
	Name:        "deleteResource",
	Description: "删除 K8s 资源",
	Parameters: jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"namespace": {
				Type:        jsonschema.String,
				Description: "资源所在的命名空间",
			},
			"resource_type": {
				Type:        jsonschema.String,
				Description: "K8s 资源标准类型，例如 Pod、Deployment、Service 等",
			},
			"resource_name": {
				Type:        jsonschema.String,
				Description: "资源名称",
			},
		},
	},
}
