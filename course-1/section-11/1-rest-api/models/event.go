package models

import (
	"restApi/db"
	"time"
)

type Event struct {
	ID          int64     `json:"id" `
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

func (e Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, dateTime, userID)
		VALUES (?, ?, ?, ?, ?)
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
		SELECT * FROM events
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `
		SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?, userID = ?
		WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error {
	query := `
		DELETE FROM events WHERE id = ?
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(e.ID)
	if err != nil {
		return err
	}

	return nil
}
