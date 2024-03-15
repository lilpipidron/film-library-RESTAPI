package film

import "time"

type Film struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release-date"`
	Rating      float32   `json:"rating"`
}
