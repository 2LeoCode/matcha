package utils

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type TypedComponent interface {
	templ.Component
	component() templ.Component
}

type TypedComponentBase struct {
	TypedComponent
}

func (this *TypedComponentBase) Render(ctx context.Context, w io.Writer) error {
	return this.component().Render(ctx, w)
}
