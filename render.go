package quarto

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
)

// Render is the interface to `quarto render`
func Render(
	ctx context.Context,
	src string,
	config *Config,
) (*Document, error) {
	cmd, err := createCommand(ctx, "render", src)
	if err != nil {
		return nil, err
	}

	config.addToCommand(cmd)

	var basename string
	outputFileName := config.Output

	// allow for path in --output based on https://github.com/quarto-dev/quarto-cli/issues/2440
	// copy the source file into the destination directory and render from there
	if outputFileName != "" {
		basename = filepath.Base(outputFileName)
		outputDir := filepath.Dir(outputFileName)
		cmd.WithArg([]string{"--output", basename})

		if basename != outputFileName {
			// Ensuring the output directory exists
			if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
				return nil, err
			}
			inputFilePath := filepath.Join(outputDir, filepath.Base(src))
			if !exist(inputFilePath) {
				// Copying the source file into the output directory
				if err := copyFile(src, inputFilePath); err != nil {
					return nil, err
				}
				defer os.Remove(inputFilePath) // removed the copied source file
			}

			cmd.WithDir(outputDir)
			cmd.WithSrc(filepath.Base(src))
		}
	}

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return &Document{src: src}, nil
}

type RenderParam struct {
	Output string
	Format string
	Params map[string]string
}

func copyFile(src, dst string) error {
	inputFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	return err
}

func exist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
