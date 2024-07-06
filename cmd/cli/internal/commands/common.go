package commands

import (
	"fmt"
	"time"
)

const TimeLayout = "2006-01-02"

func ToSubscriptionDates(start, end time.Time, duration *int) (time.Time, time.Time, error) {
	var endDate time.Time
	switch {
	case duration == nil && end.IsZero():
		return time.Time{}, time.Time{}, fmt.Errorf("either duration or end date must be provided")
	case duration == nil:
		endDate = end
	default:
		endDate = start.AddDate(0, 0, *duration)
	}

	if start.After(endDate) {
		return time.Time{}, time.Time{}, fmt.Errorf("start date cannot be after end date")
	}

	return start, endDate, nil
}
