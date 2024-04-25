package quarto

import "fmt"

// Config holds arguments to the quarto cli, such as output file, format, and dynamic parameters, other arguments should go in the `Extra` map
type Config struct {
	Output string
	Format string
	Params map[string]string
	Extra  map[string]string
}

// SetParam sets a value for a dynamic parameter
func (c *Config) SetParam(key, value string) {
	if c.Params == nil {
		c.Params = make(map[string]string)
	}
	c.Params[key] = value
}

// SetExtra sets a value for an extra argument
func (c *Config) SetExtra(key, value string) {
	if c.Extra == nil {
		c.Extra = make(map[string]string)
	}
	c.Extra[key] = value
}

func (c *Config) addToCommand(cmd *command) {
	if c.Format != "" {
		cmd.WithArg([]string{"--to", c.Format})
	}
	if c.Params != nil {
		for k, v := range c.Params {
			cmd.WithArg([]string{"-P", fmt.Sprintf("%s:%s", k, v)})
		}
	}
	if c.Extra != nil {
		for k, v := range c.Extra {
			if k != "--output" && k != "-o" {
				cmd.WithArg([]string{k, v})
			}
		}
	}
}
