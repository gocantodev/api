package app

import (
	"fmt"
	"github.com/voyago/environment"
	"gocantoserver/database/connection"
)

type Kernel struct {
	Env          environment.Env
	DBConnection *connection.Connection
}

func MakeApp(env environment.Env) (Kernel, error) {
	conn, err := connection.Make(env)

	if err != nil {
		return Kernel{}, err
	}

	return Kernel{Env: env, DBConnection: &conn}, nil
}

func (receiver Kernel) Start() {
	fmt.Println("Starting the application ...")
}
