package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func main() {
	println(uuid.NewString())
	println(uuid.New().String())
	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World!:", args[0])
		},
	}
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint)
	rootCmd.Execute()
}
