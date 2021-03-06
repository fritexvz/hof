package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/cmd/labelset"

	"github.com/hofstadter-io/hof/cmd/hof/ga"
)

var labelsetLong = `group resources, datamodels, labelsets, and more`

var LabelsetCmd = &cobra.Command{

	Use: "labelset",

	Aliases: []string{
		"L",
		"lset",
	},

	Short: "group resources, datamodels, labelsets, and more",

	Long: labelsetLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},
}

func init() {

	help := LabelsetCmd.HelpFunc()
	usage := LabelsetCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	LabelsetCmd.SetHelpFunc(thelp)
	LabelsetCmd.SetUsageFunc(tusage)

	LabelsetCmd.AddCommand(cmdlabelset.InfoCmd)
	LabelsetCmd.AddCommand(cmdlabelset.CreateCmd)
	LabelsetCmd.AddCommand(cmdlabelset.GetCmd)
	LabelsetCmd.AddCommand(cmdlabelset.SetCmd)
	LabelsetCmd.AddCommand(cmdlabelset.EditCmd)
	LabelsetCmd.AddCommand(cmdlabelset.DeleteCmd)

}
