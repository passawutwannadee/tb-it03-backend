package it03apiv1

import "github.com/passawutwannadee/tb-it03/internal/usecase/it03"

type Handler struct {
	it03 it03.UseCase
}

type Dependencies struct {
	IT03 it03.UseCase
}

func New(d Dependencies) *Handler {
	return &Handler{
		it03: d.IT03,
	}
}
