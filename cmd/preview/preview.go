/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package previewcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// previewCmd represents the preview command
var PreviewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Runs a preview of the stream at port 8000",
	Long:  `:0 Runs preview that is used to check the stream before going live.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("preview called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// previewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// previewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
