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

func TestNormalizeFunc(t *testing.T) {
	timeUTC := Now()
	timeEST, _ := timeUTC.Normalize("EST")

	if !timeUTC.Equal(timeEST) {
		t.Fatalf("Excepted times to be Equal but we got %s and %s", timeUTC, timeEST)
	}

	if timeEST.Timezone != "EST" {
		t.Fatalf("Expected the timezone to be EST but got %s instead", timeEST.Timezone)
	}
	location, _ := time.LoadLocation("EST")
	if timeEST.Location.String() != location.String() {
		t.Fatalf("Expected the location to be %s but got %s instead", location, timeEST.Location)
	}

	epochTimeUTC := timeUTC.Epoch()
	epochTimeEST := timeEST.Epoch()

	if epochTimeEST != epochTimeUTC {
		t.Fatalf("Expected epochs to match we we got %d, and %d", epochTimeUTC, epochTimeEST)
	}
}
