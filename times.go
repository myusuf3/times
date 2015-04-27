package times

import (
	"errors"
	"time"
)

type Times struct {
	time.Time
	timezone string
	location *time.Location
}

// This will return an instance of Times the current time localized to UTC
func Now() Times {
	timezone := "UTC"
	location, _ := time.LoadLocation(timezone)
	time := time.Now().UTC()
	return Times{time, timezone, location}
}

// This will localixe
func (t Times) Normalize(timezone string) (*Times, error) {
	location, err := time.LoadLocation(timezone)

	if err != nil {
		return nil, errors.New("The provided timezone isn't a valid time.Location " + timezone)
	}
	time := t.In(location)
	return &Times{time, timezone, location}, err

}
