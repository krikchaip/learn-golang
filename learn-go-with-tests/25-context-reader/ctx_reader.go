package ctx_reader

import (
	"context"
	"io"
)

// ?? "delegation pattern"
// implements: io.Reader
type ContextAwareReader struct {
	ctx      context.Context
	delegate io.Reader
}

func NewContextReader(ctx context.Context, reader io.Reader) *ContextAwareReader {
	return &ContextAwareReader{ctx, reader}
}

func (c *ContextAwareReader) Read(p []byte) (int, error) {
	if err := c.ctx.Err(); err != nil {
		return 0, err
	}

	return c.delegate.Read(p)
}
