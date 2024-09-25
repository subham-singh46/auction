package postgresDb

import (
	"errors"
	"github.com/hemantsharma1498/auction/store/models"
	"github.com/lib/pq"
)

func (pg *PostgresDb) CreateUser(user *models.User) error {
	_, err := pg.db.Exec("INSERT INTO users(name, email, salt, pw_hash, mobile) VALUES($1, $2, $3, $4, $5)", user.Name, user.Email, user.Salt, user.PwHash, user.Mobile)
	if err != nil {
		return err
	}
	return nil
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
