package models

import (
	"database/sql"
)

// User tables
type User struct {
	Name       string
	Email      string
	Registered string
}

// All register
func All(db *sql.DB) ([]*User, error) {
	rows, err := db.Query("SELECT name, email, registered FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make([]*User, 0)
	for rows.Next() {
		field := new(User)
		err := rows.Scan(&field.Name, &field.Email, &field.Registered)
		if err != nil {
			return nil, err
		}

		results = append(results, field)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// Create register
func Create(db *sql.DB, data map[string]string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(data["name"], data["email"])
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Read user
func Read(db *sql.DB, id int64) (*User, error) {
	stmt, err := db.Prepare("SELECT name, email, registered FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	field := new(User)
	err = stmt.QueryRow(id).Scan(&field.Name, &field.Email, &field.Registered)
	if err != nil {
		return nil, err
	}

	return field, nil
}

// Update users
func Update(db *sql.DB, id int, data map[string]string) error {
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(data["name"], data["email"], id)
	if err != nil {
		return err
	}

	return nil
}

// Delete user
func Delete(db *sql.DB, id int64) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
