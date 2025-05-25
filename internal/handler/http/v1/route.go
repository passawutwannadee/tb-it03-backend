package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/passawutwannadee/tb-it03/internal/app"
	it03apiv1 "github.com/passawutwannadee/tb-it03/internal/handler/http/v1/IT03"
)

func Routes(r *chi.Mux, cfg *app.AppConfig) {
	api := chi.NewRouter()

	it03apiv1.Routes(api, cfg)

	r.Mount("/api/v1", api)
}
