package times

import (
	"errors"
	"time"
)

type Times struct {
	time     time.Time
	timezone string
	location *time.Location
}

func (t *Times) String() string {
	return t.time.String()
}

// This will return an instance of Times the current time localized to UTC
func Now() *Times {
	timezone := "UTC"
	location, _ := time.LoadLocation(timezone)
	time := time.Now().UTC()
	return &Times{time, timezone, location}
}

// This will localixe
func (t *Times) Normalize(timezone string) (*Times, error) {
	location, err := time.LoadLocation(timezone)

	if err != nil {
		return nil, errors.New("The provided timezone isn't a valid time.Location " + timezone)
	}
	time := t.time.In(location)
	return &Times{time, timezone, location}, nil
}

func (t *Times) Equal(u *Times) bool {
	return t.time.Equal(u.time)
}
