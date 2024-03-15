package actor

import (
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	"time"
)

type Actor struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	Surname     string        `json:"surname"`
	Gender      gender.Gender `json:"gender"`
	DateOfBirth time.Time     `json:"date-of-birth"`
}
