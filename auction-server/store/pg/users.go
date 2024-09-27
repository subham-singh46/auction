package postgresDb

import (
	"errors"
	"fmt"

	"github.com/hemantsharma1498/auction/store/models"
	"github.com/lib/pq"
)

func (pg *PostgresDb) CreateUser(user *models.User) (int, error) {
	var userId int
	query := `INSERT INTO users(name, email, salt, pw_hash, mobile) VALUES($1, $2, $3, $4, $5) RETURNING user_id`
	err := pg.db.QueryRow(query, user.Name, user.Email, user.Salt, user.PwHash, user.Mobile).Scan(&userId)
	if err != nil {
		return -1, fmt.Errorf("failed to create user: %w", err)
	}
	return userId, nil
}

func (pg *PostgresDb) GetUsersByIds(userIds []int) ([]*models.User, error) {
	users := make([]*models.User, 0)
	rows, err := pg.db.Query("SELECT user_id, name, email, salt, pw_hash, mobile FROM users WHERE user_id = ANY($1)", pq.Array(userIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	foundRows := false
	for rows.Next() {
		foundRows = true
		user := &models.User{}
		rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Salt, &user.PwHash, &user.Mobile)
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if !foundRows {
		return nil, errors.New("no users found for the provided ids")
	}
	return users, nil
}

func (pg *PostgresDb) GetUsersByEmail(emails []string) ([]*models.User, error) {
	users := make([]*models.User, 0)
	rows, err := pg.db.Query("SELECT user_id, name, email, salt, pw_hash, mobile FROM users WHERE email = ANY($1)", pq.Array(emails))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	foundRows := false
	for rows.Next() {
		foundRows = true
		user := &models.User{}
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Salt, &user.PwHash, &user.Mobile); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	if !foundRows {
		return nil, errors.New("no users found for the provided emails")
	}

	return users, nil
}

func (pg *PostgresDb) UpdatePassword(email, salt, pwHash string) error {
	fmt.Println("salt: ", salt)
	fmt.Println("hash: ", pwHash)
	fmt.Println("email: ", email)
	res, err := pg.db.Exec("UPDATE users SET salt = $1, pw_hash = $2 WHERE email = $3", salt, pwHash, email)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	fmt.Println(rowsAffected)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}
	return nil
}
