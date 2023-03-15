package datetime

import (
	"log"
	"testing"
	"time"
)

func TestInTimeSpan(t *testing.T) {
	start := time.Now()
	end := time.Now().Add(48 * time.Hour)
	in := time.Now().Add(24 * time.Hour)

	if InTimeSpan(start, end, in) {
		t.Log("Time in span")
		return
	}

	log.Fatal("error")
}
