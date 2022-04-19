package version

import (
	"fmt"

	"github.com/sika-training-example/hc/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "foo",
	Short: "Prints \"foo!\"",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Foo!\n")
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
