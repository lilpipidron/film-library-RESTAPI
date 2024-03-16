package film

import (
	"sort"
	"time"
)

type Film struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release-date"`
	Rating      float32   `json:"rating"`
}

func ByIDAsc(films []*Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].ID < films[j].ID
	})
}

func ByIDDesc(films []*Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].ID > films[j].ID
	})
}

func ByTitleAsc(films []*Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].Title < films[j].Title
	})
}

func ByTitleDesc(films []*Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].Title > films[j].Title
	})
}

func ByReleaseDateAsc(films []*Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].ReleaseDate.Before(films[j].ReleaseDate)
	})
}

func ByReleaseDateDesc(films []*Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].ReleaseDate.After(films[j].ReleaseDate)
	})
}

func ByRatingAsc(films []*Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].Rating < films[j].Rating
	})
}

func ByRatingDesc(films []*Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].Rating > films[j].Rating
	})
}
