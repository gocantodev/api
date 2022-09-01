package Repository

import (
	"github.com/gocantodev/server/app/Database"
	"github.com/gocantodev/server/app/Entity"
)

type UsersRepository struct {
	db *Database.Connection
}

func MakeUsersRepository(db *Database.Connection) UsersRepository {
	return UsersRepository{
		db: db,
	}
}

func (receiver UsersRepository) FindByUuid() Entity.User {
	return Entity.User{}
}
