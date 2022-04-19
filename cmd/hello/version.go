package version

import (
	"fmt"

	"github.com/sika-training-example/hc/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "hello",
	Short:   "Prints \"Hello World!\"",
	Aliases: []string{"v"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("Hello World!")
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
