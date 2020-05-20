package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

// cue run + hof extra
#CmdCommand: schema.#Command & {
	TBD:   "+ "
	Name:  "cmd"
	Usage: "cmd [flags] [cmd] [args]"
	Short: "Run commands from the scripting layer"
	Long:  Short
}
