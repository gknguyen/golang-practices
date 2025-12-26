package models

import "time"

type Event struct {
	ID          int       `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events []Event = []Event{}

func (e Event) Save() {
	// TODO: add it to a database

	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
