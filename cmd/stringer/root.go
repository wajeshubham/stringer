package stringer

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// To override the version variable at build-time using ldflags:
// go build -o ./dist/stringer -ldflags="-X 'github.com/wajeshubham/stringer/cmd/stringer.version=0.0.2'" main.go

var version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "stringer",
	Short:   "stringer - a simple CLI to transform and inspect strings",
	Version: version,
	Long: `
stringer - a fancy CLI
   
One can use stringer to modify or inspect strings straight from the terminal

`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
