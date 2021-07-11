package persistence

import (
	"database/sql"
	"github.com/jaedle/time-track/service/internal/model"
)

type TokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

const createTableStatement = `CREATE TABLE tokens (
    userId varchar(255),
    token varchar(255)
);`

func (r *TokenRepository) Init() error {
	_, err := r.db.Exec(createTableStatement)
	return err
}

func (r *TokenRepository) Insert(token model.Token) error {
	_, err := r.db.Exec(`INSERT INTO tokens (userId, token) VALUES (?, ?)`, token.UserId, token.Token)
	return err
}

func (r *TokenRepository) Size() (int, error) {
	query, err := r.db.Query("SELECT COUNT(*) FROM tokens")
	if err != nil {
		return -1, err
	}

	defer func() { _ = query.Close() }()

	var count int
	query.Next()
	err = query.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TokenRepository) FindForUser(userId string) ([]model.Token, error) {
	query, err := r.db.Query("SELECT userId, token FROM tokens WHERE userId = ?", userId)
	if err != nil {
		return nil, err
	}
	defer func() { _ = query.Close() }()

	var result []model.Token
	for query.Next() {
		token := model.Token{}
		if err := query.Scan(&token.UserId, &token.Token); err != nil {
			return nil, err
		}
		result = append(result, token)
	}

	return result, nil
}
