package models

import (
	"errors"
	"todo/database"
	"todo/utils"
)

type User struct {
	Id       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) SaveUser() error {
	if len(user.Password) <= 6 {
		return errors.New("password is too short")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	var existingUser User

	query := "SELECT id FROM users WHERE email = $1"
	err = database.DB.QueryRow(query, user.Email).Scan(&existingUser.Id)

	if err == nil {
		return errors.New("User already exists")
	}

	query = "INSERT INTO users (email, secured_password, password) VALUES ($1, $2, $3) RETURNING id"

	err = database.DB.QueryRow(query, user.Email, hashedPassword, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) GetUserByEmail() error {
	query := "SELECT id, secured_password, password, email FROM users WHERE email = $1"
	row := database.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.Id, &retrievedPassword, &user.Password, &user.Email)
	if err != nil {
		return err
	}

	passwordValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !passwordValid {
		return errors.New("invalid password")
	}

	return nil
}
