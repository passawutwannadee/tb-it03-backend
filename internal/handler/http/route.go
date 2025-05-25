package httphandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/passawutwannadee/tb-it03/internal/app"
	v1 "github.com/passawutwannadee/tb-it03/internal/handler/http/v1"
)

func Routes(r *chi.Mux, cfg *app.AppConfig) {
	v1.Routes(r, cfg)
}
