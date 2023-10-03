package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {

}

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello is a command line tool",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
