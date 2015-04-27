package times

import (
	"testing"
	"time"
)

func TestNowFunc(t *testing.T) {
	timeNow := time.Now().UTC().Truncate(time.Millisecond)
	timesNow := Now()
	timesNowTime := timesNow.Time.Truncate(time.Millisecond)
	if timeNow != timesNowTime {
		t.Fatalf("Excepted %s and %s to be equal", timesNowTime, timeNow)
	}

	if timesNow.Timezone != "UTC" {
		t.Fatalf("Expected the timezone to be UTC but got %s instead", timesNow.Timezone)
	}
	location, _ := time.LoadLocation("UTC")
	if timesNow.Location != location {
		t.Fatalf("Expected the location to be %s but got %s instead", location, timesNow.Location)
	}
}
