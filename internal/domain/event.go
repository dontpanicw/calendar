package domain

import "time"

type Event struct {
	eventId     int64
	userId      int64
	date        time.Time
	description string
}
