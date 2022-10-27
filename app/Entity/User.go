package Entity

import "database/sql"

type User struct {
	Id         int64
	Uuid       string
	FirstName  string
	LastName   string
	Email      string
	Password   string //hash value
	VerifiedAt sql.NullTime
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  sql.NullTime
}
