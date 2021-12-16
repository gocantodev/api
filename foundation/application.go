package foundation

import (
	"fmt"
	"github.com/voyago/environment"
	"gocantoserver/database/connection"
)

type Application struct {
	Env          environment.Env
	DBConnection *connection.Connection
}

func CreateApp(env environment.Env) (Application, error) {
	conn, err := connection.Make(env)

	if err != nil {
		return Application{}, err
	}

	return Application{Env: env, DBConnection: &conn}, nil
}

func (receiver Application) Start() {
	fmt.Println("Starting the application ...")
}

func (receiver Application) GetConnection() *connection.Connection {
	return receiver.DBConnection
}
