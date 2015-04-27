package times

import "time"

type Times struct {
	time.Time
	timezone string
	location time.Location
}

// This will return an instance of Times the current time localized to UTC
func Now() Times {
	timezone = "UTC"
	location, err := LoadLocation(timezone)
	now := time.Now().UTC()
	return Times{now, timezone, location}
}

// This will localixe
func (t Times) Localize() {

}
