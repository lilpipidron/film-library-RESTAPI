package actor

import (
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	"time"
)

type Actor struct {
	ID          int64
	Name        string
	Surname     string
	Gender      gender.Gender
	DateOfBirth time.Time
}
