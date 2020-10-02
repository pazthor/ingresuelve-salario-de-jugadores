package main

import (
	"./cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "ingResuelve-cli"}
	rootCmd.AddCommand(cmd.InitResuelveCmd())
	rootCmd.Execute()
}
