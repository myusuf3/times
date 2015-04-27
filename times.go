package times

import "time"

type Times struct {
	time     time.Time
	timezone string
	location time.Location
}

// This will return an instance of Times the current time localized to UTC
func Now() Times {

	return Times{}
}

// This will localixe
func (t Times) Localize() {

}
