package database

import (
	"authForum/models"
	"database/sql"
	"fmt"
)

func SqlCreateUser(name, email, pass string) error {
	query := "INSERT INTO users (username, useremail, passwordhash) VALUES ($1, $2, $3)"

	_, err := DB.Exec(query, name, email, pass)
	if err != nil {
		return err
	}

	return nil
}

func SqlFindUser(email string, User *models.User) error {

	findQuery := "SELECT * FROM users WHERE useremail = $1"
	row := DB.QueryRow(findQuery, email)

	err := row.Scan(&User.Id, &User.UserName, &User.UserEmail, &User.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return err
	}

	return nil
}
