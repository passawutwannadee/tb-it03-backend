package it03apiv1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/passawutwannadee/tb-it03/internal/app"
	webv1 "github.com/passawutwannadee/tb-it03/internal/handler/http/v1/web"
	postgresrepo "github.com/passawutwannadee/tb-it03/internal/repo/postgres"
	"github.com/passawutwannadee/tb-it03/internal/util"
)

type PatchReq struct {
	IDs      []int32 `json:"ids" validate:"required"`
	Reason   string  `json:"reason" validate:"required"`
	StatusID int16   `json:"status_id" validate:"required"`
}

type PatchRes struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Reason   string `json:"reason"`
	StatusID int16  `json:"status_id"`
	Status   string `json:"status"`
}

func toPatch(req PatchReq) (*postgresrepo.IT03UpdateParams, error) {
	if req.StatusID == 1 {
		return nil, errors.New(ErrWaitingStatusNotAllow)
	}

	var filterID int16 = 1

	return &postgresrepo.IT03UpdateParams{
		Column1:      req.IDs,
		Reason:       &req.Reason,
		StatusID:     &req.StatusID,
		WithStatusID: &filterID,
	}, nil

}

func toPatchResponse(req []postgresrepo.IT03UpdateRow) *[]PatchRes {

	var res []PatchRes
	for _, v := range req {

		toAppend := PatchRes{
			ID:       v.ID,
			Name:     v.Name,
			Reason:   v.Reason,
			StatusID: v.StatusID,
		}

		if v.Status != nil {
			toAppend.Status = *v.Status
		}

		res = append(res, toAppend)
	}

	return &res
}

func (h Handler) Patch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PatchReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		webv1.Error(w, http.StatusBadRequest, app.ErrBadReq, err)
		return
	}

	if err := validator.New().Struct(req); err != nil {
		webv1.Error(w, http.StatusBadRequest, app.ErrBadReq, err)
		return
	}

	model, err := toPatch(req)
	if err != nil {
		webv1.Error(w, http.StatusBadRequest, app.ErrBadReq, err)
		return
	}

	res, err := h.it03.Update(ctx, model)
	if err != nil {
		webv1.Error(w, http.StatusInternalServerError, app.ErrInternal, err)
		return
	}

	if len(res) == 0 {
		webv1.Error(w, http.StatusNotFound, app.ErrNotFound, err)
		return
	}

	var updatedIds []int32
	for _, v := range res {
		updatedIds = append(updatedIds, v.ID)
	}

	if util.ArraysMatchInt32(req.IDs, updatedIds) {
		webv1.Respond(w, http.StatusOK, toPatchResponse(res))
	} else {
		webv1.Respond(w, http.StatusMultiStatus, toPatchResponse(res))
	}
}
