package Repository

import (
	"github.com/gocantodev/server/app/Database"
	"github.com/gocantodev/server/app/Entity"
)

type UsersRepository struct {
	connection *Database.Connection
}

func MakeUsersRepository(connection *Database.Connection) UsersRepository {
	return UsersRepository{
		connection: connection,
	}
}

func (receiver UsersRepository) FindByUuid(uuid string) (Entity.User, error) {

	user := Entity.User{}

	sql := `SELECT * FROM users WHERE uuid = $1`

	result := receiver.connection.GetDB().
		QueryRow(sql, uuid).
		Scan(
			&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email, &user.Password,
			&user.VerifiedAt, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
		)

	if result != nil {
		return Entity.User{}, result
	}

	return user, nil
}
