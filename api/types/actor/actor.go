package actor

import (
	"time"

	"github.com/lilpipidron/vk-godeveloper-task/api/types/film"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
)

type Actor struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Surname     string        `json:"surname"`
	Gender      gender.Gender `json:"gender"`
	DateOfBirth time.Time     `json:"date-of-birth"`
}

type ActorWithFilms struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Surname     string        `json:"surname"`
	Gender      gender.Gender `json:"gender"`
	DateOfBirth time.Time     `json:"date-of-birth"`
	Films       []film.Film
}
