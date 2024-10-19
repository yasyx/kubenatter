// Package cmd /*
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	clientGo "github.com/yasyx/kubenatter/pkg/client-go"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with kubernetes.",
	Long: `Chat with kubernetes by diff AI. For example:
	The command "kubenatter chat -m gpt4o" will chat with kubernetes using gpt4o model.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("chat with model: %s，kubeconfig location: %s\n", model, kubeConfig)

		clientGo, err := clientGo.NewClient(kubeConfig)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		podList, err := clientGo.ClientSet.CoreV1().Pods("").List(context.Background(), metaV1.ListOptions{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, p := range podList.Items {
			fmt.Println(p.Name)
		}

		gvr := schema.GroupVersionResource{
			Group:    "",
			Version:  "v1",
			Resource: "pods",
		}

		unstructuredObj, err := clientGo.DynamicClient.Resource(gvr).List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		podList = &coreV1.PodList{}

		err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj.UnstructuredContent(), podList)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, pod := range podList.Items {
			fmt.Println(pod.Name)
		}

	},
}

var model string

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringVarP(&model, "model", "m", "gpt4o", "model to use for chat")
}
