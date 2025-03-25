package validationtool

import (
	"fmt"
	"time"
)

func ValidateDate(value string, name string) error {

	// check if value is a valid date or not if not return error
	if _, err := time.Parse("2006-01-02", value); err != nil {

		return fmt.Errorf("invalid date for %s", name)
	}

	return nil
}

func ValidateTime(value string, name string) error {

	// check valu is time and formatted like this 10:01:00
	if _, err := time.Parse("15:04:05", value); err != nil {

		return fmt.Errorf("invalid time for %s", name)
	}

	return nil
}
