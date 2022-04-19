
package root

import (
	"github.com/sika-training-example/hc/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "hc",
	Short: "hc, " + version.Version,
}
