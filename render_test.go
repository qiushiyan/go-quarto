package quarto

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	ctx := context.Background()
	tmpl, err := os.Open("assets/hello.qmd")
	if err != nil {
		require.NoError(t, err)
	}
	defer tmpl.Close()

	cases := []struct {
		name        string
		createParam func() *Config
	}{
		{

			name: "ok",
			createParam: func() *Config {
				c := &Config{
					Output: os.TempDir() + "test.html",
					Format: "html",
				}

				c.SetParam("n", "200")
				return c
			},
		},
	}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.CreateTemp("", "*.qmd")
			require.NoError(t, err)
			defer os.Remove(f.Name())

			io.Copy(f, tmpl)
			_, err = Render(ctx, f.Name(), tc.createParam())
			require.NoError(t, err)
		})

	}
}
