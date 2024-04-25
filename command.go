package quarto

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
)

type command struct {
	c      *exec.Cmd
	errBuf *bytes.Buffer
}

func createCommand(ctx context.Context, action, src string) (*command, error) {
	q := findQuarto()
	if q == "" {
		return nil, ErrQuartoNotFound
	}

	errBuf := new(bytes.Buffer)
	cmd := exec.CommandContext(ctx, q, action, src)
	cmd.Stderr = errBuf

	return &command{
		c:      cmd,
		errBuf: errBuf,
	}, nil
}

func (c *command) WithSrc(src string) {
	c.c.Args[2] = src
}

func (c *command) WithDir(dir string) {
	c.c.Dir = dir
}

func (c *command) WithArg(arg []string) {
	c.c.Args = append(c.c.Args, arg...)
}

func (c *command) WithArgMap(arg map[string]string) {
	for k, v := range arg {
		c.c.Args = append(c.c.Args, k, v)
	}
}

func (c *command) Run() error {
	if err := c.c.Run(); err != nil {
		return errors.New(c.errBuf.String())
	}

	return nil
}
