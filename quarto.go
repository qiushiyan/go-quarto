package quarto

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"strings"
)

// ErrQuartoNotFound is returned when Quarto is not found in the PATH or QUARTO_PATH environment variable.
var ErrQuartoNotFound = errors.New(
	"Quarto command-line tools path not found! Please make sure you have installed and added Quarto to your PATH or set the QUARTO_PATH environment variable",
)

func Run(ctx context.Context, arg ...string) (*Document, error) {
	q := findQuarto()
	if q == "" {
		return &Document{}, ErrQuartoNotFound
	}

	cmd := exec.CommandContext(ctx, q, arg...)

	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	// var outBuf bytes.Buffer

	// cmd.Stdout = &outBuf
	if err := cmd.Run(); err != nil {
		return &Document{}, errors.New(errBuf.String())
	}

	return &Document{path: arg[1]}, nil
}

func Render(ctx context.Context, arg ...string) (*Document, error) {
	s := append([]string{"render"}, arg...)
	return Run(ctx, s...)
}

func findQuarto() string {
	pathEnv := os.Getenv("QUARTO_PATH")
	if pathEnv == "" {
		cmd := exec.Command("which", "quarto")
		output, err := cmd.Output()
		if err != nil {
			return ""
		}

		val := strings.TrimSpace(string(output))
		return val
	}

	return pathEnv

}
