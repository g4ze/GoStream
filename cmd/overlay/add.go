/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package overlaycmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "ads the new overlay to the stream",
	Long:  `adds the new overlay to the stream.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gostream overlay added file %s to overlays", filepath)

	},
}
var (
	filepath string
)

func init() {
	OverlayCmd.AddCommand(AddCmd)

	AddCmd.Flags().StringVarP(&filepath, "file", "f", "", "file path of the overlay")

	if err := AddCmd.MarkFlagRequired("file"); err != nil {
		fmt.Println(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
