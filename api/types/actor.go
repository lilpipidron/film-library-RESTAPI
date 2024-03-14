package types

import "time"

type Actor struct {
	ID          int64
	Name        string
	Surname     string
	Gender      Gender
	DateOfBirth time.Time
}
