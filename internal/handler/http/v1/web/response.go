package webv1

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/passawutwannadee/tb-it03/config"
	"github.com/passawutwannadee/tb-it03/internal/util"
)

type AppResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Errors struct {
	Reason string `json:"reason,omitempty"` // Reason is a machine-readable description of the error.
}

type Err struct {
	Code    int     `json:"code"`
	Message string  `json:"message"` // Message is a human-readable description of the error.
	Errors  *Errors `json:"errors,omitempty"`
}

type AppError struct {
	Success bool `json:"success"`
	Error   Err  `json:"error"`
}

type Meta struct {
	PaginationType string      `json:"pagination_type"`
	Pagination     interface{} `json:"pagination"`
}

type PaginationCursor struct {
	Limit  int64  `json:"limit"`
	Cursor string `json:"cursor"`
}

type PaginationOffset struct {
	CurrentPage     int32 `json:"current_page"`
	PageSize        int32 `json:"page_size"`
	HasNextPage     bool  `json:"has_next_page"`
	HasPreviousPage bool  `json:"has_previous_page"`
}

const (
	Cursor = "cursor"
	Offset = "offset"
)

// Respond is a utility function to send any response with a given status code.
func respond(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)
}

// RespondSuccess sends a standard success response.
func Respond(w http.ResponseWriter, statusCode int, data interface{}) {
	respond(w, statusCode, AppResponse{
		Success: true,
		Data:    data,
	})
}

func OffsetResponse(w http.ResponseWriter, data interface{}, totalItems int64, page int32, pageSize int32) {
	page64, _ := util.Int32ToInt64(page)
	pageSize64, _ := util.Int32ToInt64(pageSize)

	var totalPages int64
	if pageSize == 0 {
		totalPages = 0
	} else {
		totalPages = (100 + pageSize64 - 1) / pageSize64
	}
	hasNextPage := page64 < totalPages
	hasPreviousPage := page > 1

	respond(w, http.StatusOK, AppResponse{
		Success: true,
		Data:    data,
		Meta: &Meta{
			PaginationType: Offset,
			Pagination: &PaginationOffset{
				CurrentPage:     page,
				PageSize:        pageSize,
				HasNextPage:     hasNextPage,
				HasPreviousPage: hasPreviousPage,
			},
		},
	})
}

func CursorResponse(w http.ResponseWriter, data interface{}, cursor *string, limit int64) {

	res := AppResponse{
		Success: true,
		Data:    data,
		Meta: &Meta{
			PaginationType: Cursor,
			Pagination: &PaginationCursor{
				Limit:  limit,
				Cursor: "",
			},
		},
	}

	if cursor != nil {

		encoded := base64.StdEncoding.EncodeToString([]byte(*cursor))
		res.Meta.Pagination.(*PaginationCursor).Cursor = encoded
	}

	respond(w, http.StatusOK, res)
}

// RespondError sends a standard error response.
func Error(w http.ResponseWriter, code int, message string, reason error) {
	err := AppError{
		Success: false,
		Error: Err{
			Code:    code,
			Message: message,
		},
	}
	if config.C.AppMode == config.AppDebug {
		if reason != nil {
			err.Error.Errors = &Errors{
				Reason: reason.Error(),
			}
		}
	}

	respond(w, code, err)
}
