package utils

import (
	"errors"
	"time"
)

const layout = "2006-01-02"

func IsTaskDateStringValid(date string) error {
	taskDate, err := GetFormattedDateFromString(date)
	if err != nil {
		return err
	}

	currentDate := time.Now().Truncate(24 * time.Hour)

	if taskDate.Before(currentDate) {
		return errors.New("task date should be equal or greater then current date")
	}

	return nil
}

func GetFormattedDateFromString(date string) (time.Time, error) {
	return time.Parse(layout, date)
}
