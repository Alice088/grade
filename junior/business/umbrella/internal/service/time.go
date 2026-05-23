package service

import (
	"fmt"
	"time"
)

type Time struct {
	StartDate time.Time
}

func NewTime(startDate time.Time) Time {
	return Time{
		StartDate: startDate,
	}
}

func (t *Time) Since() string {
	now := time.Now().UTC()
	left := now.UTC().Sub(t.StartDate)

	days := int(left.Hours()) / 24
	hours := int(left.Hours()) % 24
	minutes := int(left.Minutes()) % 60

	return fmt.Sprintf("%dD-%dH-%dM", days, hours, minutes)
}
