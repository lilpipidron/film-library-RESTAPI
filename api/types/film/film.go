package film

import "time"

type film struct {
	ID          int64
	Title       string
	Description string
	ReleaseDate time.Time
	Rating      float32
}
