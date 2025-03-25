package validationtool

import (
	"fmt"
)

func StringMinAndMaxLength(value string, name string, min int, max int) error {

	if err := StringMinLength(value, name, min); err != nil {

		return err
	}

	return StringMaxLength(value, name, max)
}

func StringMinLength(value string, name string, length int) error {

	if len(value) < length {

		return fmt.Errorf("min length for %s is %d", name, length)
	}

	return nil
}

func StringMaxLength(value string, name string, length int) error {

	if len(value) > length {

		return fmt.Errorf("max length for %s is %d", name, length)
	}

	return nil
}
