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

func Version() (string, error) {
	quarto := findQuarto()
	if quarto == "" {
		return "", ErrQuartoNotFound
	}

	cmd := exec.Command(quarto, "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func findQuarto() string {
	p := os.Getenv("QUARTO_PATH")
	if p == "" {
		cmd := exec.Command("which", "quarto")
		output, err := cmd.Output()
		if err != nil {
			return ""
		}

		val := strings.TrimSpace(string(output))
		return val
	}

	return p

}
