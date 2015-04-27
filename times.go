package times

import (
	"errors"
	"time"
)

type Times struct {
	Time     time.Time
	Timezone string
	Location *time.Location
}

func (t *Times) String() string {
	return t.Time.String()
}

// This will return an instance of Times the current time localized to UTC
func Now() *Times {
	timezone := "UTC"
	location, _ := time.LoadLocation(timezone)
	time := time.Now().UTC()
	return &Times{time, timezone, location}
}

func EpochNow() int64 {
	return Now().Epoch()
}

// This will localixe
func (t *Times) Normalize(timezone string) (*Times, error) {
	location, err := time.LoadLocation(timezone)

	if err != nil {
		return nil, errors.New("The provided timezone isn't a valid time.Location " + timezone)
	}
	time := t.Time.In(location)
	return &Times{time, timezone, location}, nil
}

func (t *Times) Equal(u *Times) bool {
	return t.Time.Equal(u.Time)
}

func (t *Times) Epoch() int64 {
	return t.Time.Unix()
}
