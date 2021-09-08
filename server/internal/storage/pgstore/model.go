package pgstore

import "time"

type User struct {
	ID        uint64
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
}

type UserList []User
