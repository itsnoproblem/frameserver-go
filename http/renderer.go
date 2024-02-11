package http

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
)

var DefaultRenderer = &renderer{}

func NewRenderer(ctx context.Context, status int, component templ.Component) *renderer {
	return &renderer{
		ctx:       ctx,
		Status:    status,
		Component: component,
	}
}

type renderer struct {
	ctx       context.Context
	Status    int
	Component templ.Component
}

func (t renderer) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Status)
	if t.Component != nil {
		return t.Component.Render(t.ctx, w)
	}
	return nil
}

func (t renderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *renderer) Instance(name string, data any) render.Render {
	templData, ok := data.(templ.Component)
	if !ok {
		return nil
	}
	return &renderer{
		ctx:       context.Background(),
		Status:    http.StatusOK,
		Component: templData,
	}
}
