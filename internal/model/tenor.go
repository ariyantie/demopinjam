package models

import "time"

// Tenor adalah model struct untuk tabel tenor
type Tenor struct {
	ID            int        `json:"id"`
	Tenor         int        `json:"tenor"`
	Value         float64    `json:"value"`
	PossibleLimit *float64   `json:"possible_limit,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}
