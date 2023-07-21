/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package streamcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var StreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "streams the video to the server",
	Long:  `streams video to the video streaming service with the given key`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stream called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// streamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// streamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
