package models

import (
	"restApi/db"
	"restApi/utils"
)

type Users struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *Users) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
