package stringer

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wajeshubham/stringer/pkg/stringer"
)

var countDigits = false
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Args:    cobra.ExactArgs(1),
	Short:   "Inspects a string",
	Run: func(c *cobra.Command, args []string) {
		arg := args[0]
		count, kind := stringer.Inspect(arg, countDigits)
		plural := "s"
		if count == 1 {
			plural = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", arg, count, kind, plural)
	},
}

func init() {
	inspectCmd.Flags().BoolVarP(&countDigits, "digits", "d", countDigits, "Counts only digits")
	rootCmd.AddCommand(inspectCmd)
}
