package times

import "time"

struct Times {
    time time.Time
    timezone string
    location Location
}

// This will return an instance of Times the current time localized to UTC
func Now() times.Times{

	return Times{}
}

// This will localixe
func (t Time) Localize( )