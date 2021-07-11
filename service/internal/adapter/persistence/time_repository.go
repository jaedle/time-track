package persistence

import (
	"database/sql"
	"fmt"
)

type TimeRepository struct {
	db *sql.DB
}

func NewTimeRepository(db *sql.DB) *TimeRepository {
	return &TimeRepository{
		db: db,
	}
}

func (r *TimeRepository) Init() error {
	_, err := r.db.Exec(`CREATE TABLE times (
    description varchar(255)
);`)
	return err
}

func (r *TimeRepository) Insert(s string) error {
	_, err := r.db.Exec(fmt.Sprintf(`INSERT INTO times (description) VALUES ("%s")`, s))
	return err
}

func (r *TimeRepository) Size() (int, error) {
	query, err := r.db.Query("SELECT COUNT(description) FROM times")
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
