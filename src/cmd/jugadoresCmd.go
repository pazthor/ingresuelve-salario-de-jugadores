package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

// InitResuelveCmd initialize beers command
func InitResuelveCmd() *cobra.Command {
	retoResuelveCmd := &cobra.Command{
		Use:   "calc-salario",
		Short: "Calcular salario de jugadores",
		Run:   runBeersFn(),
	}

	return retoResuelveCmd
}

func runBeersFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
	}
}
