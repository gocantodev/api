package Entity

import "time"

type User struct {
	Id         int64
	Uuid       string
	FirstName  string
	LastName   string
	email      string
	password   string
	VerifiedAt time.Time
	CreatedAt  time.Time
	DeletedAt  time.Time
}
