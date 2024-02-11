package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/itsnoproblem/frameserver-go/templates"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

type HandlerFunc func(ctx context.Context, request interface{}) (response interface{}, err error)
type DecoderFunc func(ctx context.Context, request *http.Request) (decoded interface{}, err error)
type FormatterFunc func(ctx context.Context, response interface{}) (templ.Component, error)

type HandlerMaker func(decoder DecoderFunc, handler HandlerFunc, formatter FormatterFunc) gin.HandlerFunc

func MakeHandler(decoder DecoderFunc, handler HandlerFunc, formatter FormatterFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := ctx.Request
		w := ctx.Writer

		decoded, err := decoder(ctx, r)
		if err != nil {
			renderError(ctx, w, err, http.StatusInternalServerError)
			return
		}

		res, err := handler(ctx, decoded)
		if err != nil {
			renderError(ctx, w, err, http.StatusInternalServerError)
			return
		}

		cmp, err := formatter(ctx, res)
		if err != nil {
			renderError(ctx, w, err, http.StatusInternalServerError)
			return
		}

		rnd := NewRenderer(ctx, http.StatusOK, cmp)
		if err = rnd.Render(w); err != nil {
			renderError(ctx, w, err, http.StatusInternalServerError)
		}
	}
}

func renderError(ctx context.Context, w http.ResponseWriter, err error, code int) {
	view := templates.ErrorView{
		Code:    code,
		Message: err.Error(),
	}

	log.Printf(">> ERROR << %s", err.Error())
	cmp := templates.ErrorFrame(view)
	_ = NewRenderer(ctx, code, cmp).Render(w)
}
