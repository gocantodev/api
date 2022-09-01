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

	db := receiver.connection.GetDB()

	row := db.QueryRow(sql, uuid)

	err := row.Scan(
		&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email, &user.Password,
		&user.VerifiedAt, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)

	if err != nil {
		return Entity.User{}, err
	}

	return user, nil
}
