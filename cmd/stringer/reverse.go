package stringer

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wajeshubham/stringer/pkg/stringer"
)

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverses a provided strong",
	Args:    cobra.ExactArgs(1), // ExactArgs returns an error if there are not exactly n args.
	Run: func(c *cobra.Command, args []string) {
		res := stringer.Reverse(args[0])
		fmt.Printf("Reveres of %s is %s\n", args[0], res)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
