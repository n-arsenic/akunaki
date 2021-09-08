package pgstore

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	conn *pgxpool.Pool
}

func Open(url string) (*Store, error) {
	conn, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &Store{conn: conn}, nil
}

func (s *Store) Close() {
	s.conn.Close()
}

func (s *Store) CreateUser(ctx context.Context, user User) (uint64, error) {
	row := s.conn.QueryRow(ctx, "INSERT INTO users (name, password, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Password, user.Email, time.Now())
	var uid uint64
	err := row.Scan(&uid)

	return uid, err
}

func (s *Store) DeleteUser(ctx context.Context, userID uint64) error {
	_, err := s.conn.Exec(ctx, "DELETE FROM users WHERE id = $1", userID)
	return err
}
