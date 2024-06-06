package models

import "time"

type Perlengkapan struct {
    ID         int       `json:"id"`
    Nama       string    `json:"nama"`
    CreatedAt  time.Time `json:"created_at"`
}
