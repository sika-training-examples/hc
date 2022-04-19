package version

import (
	"fmt"

	"github.com/sika-training-example/hc/cmd/root"
	"github.com/spf13/cobra"
)

var FlagName = ""

var Cmd = &cobra.Command{
	Use:     "hello",
	Short:   "Prints \"Hello World!\"",
	Aliases: []string{"v"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Hello %s!\n", FlagName)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Hello to someone",
	)
}
