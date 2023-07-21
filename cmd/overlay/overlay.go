/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package overlaycmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// OverlayCmd represents the overlay command
var OverlayCmd = &cobra.Command{
	Use:   "overlay",
	Short: "attach overlay for stream",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("overlay called")

		//
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overlayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overlayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
