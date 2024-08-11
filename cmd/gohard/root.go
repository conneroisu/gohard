// Package hardgo is a file format and lightweight Go library for working
// with hardware description langages.
package main

import (
	"os"

	"github.com/conneroisu/hardgo/internal/graph"
	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hardgo",
	Short: "A utility for the gohard language.",
	Long: `
Hardgo is a utility for the gohard language.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	graph.Execute()

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hardgo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// n, err := sitter.ParseCtx(
	//         context.Background(),
	//         []byte(code),
	//         verilog.GetLanguage(),
	// )
	// if err != nil {
	//         fmt.Printf("Expected no error got %v", err)
	//         os.Exit(1)
	//
	// }

	// if act := n.String(); act != expected {
	//         fmt.Printf("Expected %q got %q", expected, act)
	//         os.Exit(1)
	// }
}
