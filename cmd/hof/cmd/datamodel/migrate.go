package cmddatamodel

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/hof/cmd/hof/ga"

	"github.com/hofstadter-io/hof/lib/datamodel"
)

var migrateLong = `calculate a changeset for a data model`

func MigrateRun(args []string) (err error) {

	// you can safely comment this print out
	// fmt.Println("not implemented")

	err = datamodel.RunMigrateFromArgs(args)

	return err
}

var MigrateCmd = &cobra.Command{

	Use: "migrate",

	Aliases: []string{
		"mig",
		"migs",
		"migrations",
	},

	Short: "calculate a changeset for a data model",

	Long: migrateLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendCommandPath(cmd.CommandPath())

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = MigrateRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {

	help := MigrateCmd.HelpFunc()
	usage := MigrateCmd.UsageFunc()

	thelp := func(cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func(cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
		return usage(cmd)
	}
	MigrateCmd.SetHelpFunc(thelp)
	MigrateCmd.SetUsageFunc(tusage)

}