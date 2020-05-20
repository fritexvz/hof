package cmds

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#AuthCommand: schema.#Command & {
	TBD:   "+ "
	Name:  "auth"
	Usage: "auth"
	Short: "authentication subcommands"
	Long:  Short

	OmitRun: true

	Commands: [
		{
			TBD:   "+ "
			Name:  "login"
			Usage: "login <where>"
			Short: "login to an account, provider, system, or url"
			Long:  Short
			Args: [
				{
					Name: "where"
					Type: "string"
					Help: "the account, provider, system, or url you wish to authenticate with"
				},
			]
		},
		{
			TBD:   "+ "
			Name:  "logout"
			Usage: "logout <name>"
			Short: "logout of an authenticated session"
			Long:  Short
		},
		{
			TBD:   "+ "
			Name:  "list"
			Usage: "list"
			Short: "list known auth configurations and sessions"
			Long:  Short
		},
		{
			TBD:   "+ "
			Name:  "test"
			Usage: "test [name]"
			Short: "test your auth configuration, defaults to current context"
			Long:  Short
		},
	]
}

#ContextCommand: schema.#Command & {
	TBD:   "+ "
	Name:  "context"
	Usage: "context"
	Short: "Get, set, and use contexts"
	Long:  Short

	OmitRun: true

	Commands: [
		{
			TBD:   "+ "
			Name:  "get"
			Usage: "get <all> | <name>"
			Short: "print a context, defauts to current"
			Long:  Short
		},
		{
			TBD:   "+ "
			Name:  "set"
			Usage: "set [key.path] [value]"
			Short: "set context value, defaults to current"
			Long:  Short
		},
		{
			TBD:   "+ "
			Name:  "use"
			Usage: "use [name]"
			Short: "set a context as the current default"
			Long:  Short
		},
	]
}

#ConfigCommand: schema.#Command & {
	Name:  "config"
	Usage: "config"
	Short: "Manage local configurations"
	Long:  Short

	OmitRun: true

	Commands: [
		{
			Name:  "get"
			Usage: "get <key.path>"
			Short: "print a config or value(s) at path(s)"
			Long:  Short
		},
		{
			TBD:   "+ "
			Name:  "set"
			Usage: "set [expr]"
			Short: "set config values with expr"
			Long:  Short
			Args: [{
				Name:     "expr"
				Type:     "string"
				Required: true
				Help:     "Cue expr for value you'd like to merge into your config"
			}]
		},
	]
}

#SecretCommand: schema.#Command & {
	Name:  "secret"
	Usage: "secret"
	Short: "Manage local secrets"
	Long:  Short

	OmitRun: true

	Commands: [
		{
			Name:  "get"
			Usage: "get <key.path>"
			Short: "print a secret or value(s) at path(s)"
			Long:  Short
		},
		{
			TBD:   "+ "
			Name:  "set"
			Usage: "set [key.path] [value]"
			Short: "set secret value at path"
			Long:  Short
			Args: [{
				Name:     "path"
				Type:     "string"
				Required: true
				Help:     "Cue path to the value you'd like, if none, everything is printed"
			},{
				Name:     "value"
				Type:     "string"
				Required: true
				Help:     "Cue value is a string, or use @filename.cue"
			}]
		},
	]
}
