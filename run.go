package quarto

import "context"

// Run is the general interface to run a Quarto command with parameters
func Run(
	ctx context.Context,
	action string,
	src string,
	config *Config,
) (*Document, error) {
	cmd, err := createCommand(ctx, action, src)
	if err != nil {
		return nil, err
	}

	config.addToCommand(cmd)

	// cmd.Stdout = &outBuf
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return &Document{src: src}, nil
}
