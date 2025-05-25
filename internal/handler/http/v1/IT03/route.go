package it03apiv1

import (
	"github.com/go-chi/chi/v5"
	"github.com/passawutwannadee/tb-it03/internal/app"
	"github.com/passawutwannadee/tb-it03/internal/usecase/it03"
)

func Routes(r chi.Router, cfg *app.AppConfig) {

	it03 := it03.New(cfg.Database)
	h := New(Dependencies{
		IT03: it03,
	})

	//group for media
	r.Route("/IT03", func(r chi.Router) {
		r.Get("/", h.List)
		r.Patch("/", h.Patch)
	})

}
