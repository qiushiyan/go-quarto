package quarto

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// ErrQuartoNotFound is returned when Quarto is not found in the PATH or QUARTO_PATH environment variable.
var ErrQuartoNotFound = errors.New(
	"Quarto command-line tools path not found! Please make sure you have installed and added Quarto to your PATH or set the QUARTO_PATH environment variable",
)

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
