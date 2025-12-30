package models

import (
	"errors"
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

func (user *Users) ValidateCredentials() error {
	query := `
		SELECT id, password
		FROM users
		WHERE email = ?
	`

	var storedHashedPassword string
	err := db.DB.QueryRow(query, user.Email).Scan(&user.ID, &storedHashedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	isValidPassword := utils.CheckPasswordHash(user.Password, storedHashedPassword)
	if !isValidPassword {
		return errors.New("Credentials invalid")
	}

	return nil
}
