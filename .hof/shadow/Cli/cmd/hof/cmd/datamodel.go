package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/cmd/datamodel"

	"github.com/hofstadter-io/hof/cmd/hof/ga"
)

var datamodelLong = `create, view, diff, calculate / migrate, and manage your data models`

var DatamodelCmd = &cobra.Command{

	Use: "datamodel",

	Aliases: []string{
		"dmod",
	},

	Short: "create, view, diff, calculate / migrate, and manage your data models",

	Long: datamodelLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},
}

func init() {

	help := DatamodelCmd.HelpFunc()
	usage := DatamodelCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/help", "<omit>", 0)
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/usage", "<omit>", 0)
		return usage(cmd)
	}
	DatamodelCmd.SetHelpFunc(thelp)
	DatamodelCmd.SetUsageFunc(tusage)

	DatamodelCmd.AddCommand(cmddatamodel.ViewCmd)
	DatamodelCmd.AddCommand(cmddatamodel.ListCmd)
	DatamodelCmd.AddCommand(cmddatamodel.StatusCmd)
	DatamodelCmd.AddCommand(cmddatamodel.GraphCmd)
	DatamodelCmd.AddCommand(cmddatamodel.DiffCmd)
	DatamodelCmd.AddCommand(cmddatamodel.MigrateCmd)
	DatamodelCmd.AddCommand(cmddatamodel.TestCmd)
	DatamodelCmd.AddCommand(cmddatamodel.DeleteCmd)

}