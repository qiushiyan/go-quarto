package quarto

import "context"

// A document represents a Quarto document
type Document struct {
	path string
}

func (d *Document) Render(ctx context.Context, arg ...string) (*Document, error) {
	s := append([]string{d.path}, arg...)
	return Render(ctx, s...)
}
