package it03apiv1

import (
	"net/http"

	"github.com/passawutwannadee/tb-it03/internal/app"
	webv1 "github.com/passawutwannadee/tb-it03/internal/handler/http/v1/web"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	res, err := h.it03.OffsetList(c)
	if err != nil {
		webv1.Error(w, http.StatusInternalServerError, app.ErrInternal, err)
		return
	}

	webv1.Respond(w, http.StatusOK, res)
}
