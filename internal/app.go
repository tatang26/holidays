package internal

import (
	"cmp"
	"coding/session/internal/holiday"
	"embed"
	"net/http"
	"os"

	"go.leapkit.dev/core/render"
	"go.leapkit.dev/core/server"
)

//go:embed **/*.html **/*.html
var templates embed.FS

type Server interface {
	Addr() string
	Handler() http.Handler
}

func New() Server {
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
	)

	r.Use(render.Middleware(
		render.TemplateFS(templates, "internal"),
		render.WithDefaultLayout("layout.html"),
	))

	r.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		render.FromCtx(r.Context()).RenderClean("layout.html")
	})

	r.HandleFunc("GET /holidays/{$}", holiday.Handler)

	return r
}
