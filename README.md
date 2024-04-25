# Quarto

Go interface to Quarto (unofficial).

For more information on Quarto, see the project Github repositories at https://github.com/quarto-dev/

## Usage

Install the package with:

```
go get -u github.com/qiushiyan/go-quarto
```

Render a Quarto document

```go
config := &quarto.Config{
    Output: "path/to/output.pdf",
    Format: "html"
}
config.SetExtra("-M", "echo:true")

quarto.Render(ctx, "path/to/source.qmd", config)
```

### Specifying output

Currently Quarto does not allow the `--output` argument to contain a directory path (see [discussion](https://github.com/quarto-dev/quarto-cli/issues/2440)), this library makes a workaround to copy the source file into the specified directory and render from there. This may require adjustments to the source file as the working directory at render time is now the output directory.
