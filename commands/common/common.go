package common

import (
	"github.com/spf13/cobra"
)

// Command implements common command functions
type Command struct {
	Args []*string
}

// AddArg makes the command aware that it is expecting an argument
// The order that args are added is important
func (c *Command) AddArg(arg *string) {
	if arg == nil {
		return
	}

	c.Args = append(c.Args, arg)
}

// ParseArgs writes the arg values to the specified locations.
func (c *Command) ParseArgs() cobra.PositionalArgs {
	return func(_ *cobra.Command, args []string) error {
		for i, arg := range c.Args {
			if len(args) < i+1 {
				return nil
			}

			*arg = args[i]
		}

		return nil
	}
}
