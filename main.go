/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	chatclient "github.com/yasyx/kubenatter/pkg/chat-client"
	"reflect"
)

func main() {

	cli := chatclient.OpenAIChatClient{}
	elem := reflect.ValueOf(cli).Type()

	for i := 0; i < elem.NumMethod(); i++ {
		method := elem.Method(i)
		fmt.Println("Method name:", method.Name)
		fmt.Println("Method type:", method.Type)
	}

	//cmd.Execute()
}
