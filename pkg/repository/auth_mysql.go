package repository

import (
	"database/sql"
	"fmt"
	"todo"
)

type AuthMysql struct {
	db *sql.DB
}

func NewAuthMysql(db *sql.DB) *AuthMysql {
	return &AuthMysql{db}
}

func (r *AuthMysql) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values (?, ?, ?);", usersTable)

	_ = r.db.QueryRow(query, user.Name, user.Username, user.Password)

	query = fmt.Sprintf("SELECT id FROM %s WHERE (`name`=?) AND (`username`=?) AND (`password_hash`=?);", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthMysql) GetUser(username, password string) (int, error) {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE (`username`=?) AND (`password_hash`=?);", usersTable)

	row := r.db.QueryRow(query, username, password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}
