// Package cmd /*
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	chatclient "github.com/yasyx/kubenatter/pkg/chat-client"
	"os"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Chat with kubernetes.",
	Long: `Chat with kubernetes by different AI. For example:
	The command "kubenatter chat -m gpt-4o" will chat with kubernetes using gpt4o model.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("You are chat with the model: %s，kubeconfig location: %s\n", model, kubeConfig)
		startChat()
		//clientGo, err := clientGo.NewClient(kubeConfig)
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//
		//podList, err := clientGo.ClientSet.CoreV1().Pods("").List(context.Background(), metaV1.ListOptions{})
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//
		//for _, p := range podList.Items {
		//	fmt.Println(p.Name)
		//}
		//
		//gvr := schema.GroupVersionResource{
		//	Group:    "",
		//	Version:  "v1",
		//	Resource: "pods",
		//}
		//
		//unstructuredObj, err := clientGo.DynamicClient.Resource(gvr).List(context.TODO(), metaV1.ListOptions{})
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//
		//podList = &coreV1.PodList{}
		//
		//err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObj.UnstructuredContent(), podList)
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//
		//for _, pod := range podList.Items {
		//	fmt.Println(pod.Name)
		//}

	},
}

func startChat() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Kubenatter! Type 'exit' to quit.")
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			fmt.Println("bye!")
			break
		}
		if text == "" {
			continue
		}
		res, err := processMessage(text)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print("< ")
		fmt.Println(res)
		fmt.Print("> ")
	}
}

func processMessage(text string) (string, error) {
	//fmt.Println("receive  message: ", text)
	chatclient := chatclient.NewChatClient(model)
	return chatclient.SendMessage(text)
}

var model string

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringVarP(&model, "model", "m", "gpt-4o-mini", "model to use for chat")
}
