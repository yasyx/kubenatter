// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubenatter",
	Short: "Have a natter with kubernetes.",
	Long:  `Have a natter with kubernetes.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var kubeConfig string
var namespace string

func init() {
	defaultKubeConfig := ""
	if home := homedir.HomeDir(); home != "" {
		defaultKubeConfig = home + "/.kube/config"
	}
	rootCmd.PersistentFlags().StringVarP(&kubeConfig, "kubeConfig", "c", defaultKubeConfig, "The path to the kubernetes config file.")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "The namespace of the kubernetes resource.")
}
