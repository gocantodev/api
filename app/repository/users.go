package repository

import (
	"github.com/gocantodev/api/app/database"
	"github.com/gocantodev/api/app/entity"
)

type UsersRepository struct {
	connection *database.Connection
}

func MakeUsersRepository(connection *database.Connection) UsersRepository {
	return UsersRepository{
		connection: connection,
	}
}

func (receiver UsersRepository) FindByUuid(uuid string) (entity.User, error) {

	user := entity.User{}

	sql := `SELECT * FROM users WHERE uuid = $1`

	result := receiver.connection.GetDB().
		QueryRow(sql, uuid).
		Scan(
			&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email, &user.Password,
			&user.VerifiedAt, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
		)

	if result != nil {
		return entity.User{}, result
	}

	return user, nil
}
