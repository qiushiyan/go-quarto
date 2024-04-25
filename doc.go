/*
Package conf provides interface to the "Quarto" markdown publishing system.

For more information on Quarto, see the project GitHub repositories at https://github.com/quarto-dev/

# Example Usage

	quarto.Render(context.Background(), "example.qmd", &Config{
					Output: "output/example.pdf",
	})
*/
package quarto
