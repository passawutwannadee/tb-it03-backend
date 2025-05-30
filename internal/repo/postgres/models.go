// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package postgresrepo

import (
	"time"
)

type It03 struct {
	ID        int32      `json:"id"`
	Name      string     `json:"name"`
	Reason    string     `json:"reason"`
	StatusID  int16      `json:"status_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type It03Status struct {
	ID     int32  `json:"id"`
	Status string `json:"status"`
}
