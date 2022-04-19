package cmd

import (
	_ "github.com/sika-training-example/hc/cmd/foo"
	_ "github.com/sika-training-example/hc/cmd/hello"
	"github.com/sika-training-example/hc/cmd/root"
	_ "github.com/sika-training-example/hc/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
