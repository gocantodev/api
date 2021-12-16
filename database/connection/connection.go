package connection

import (
	"fmt"
	"github.com/voyago/environment"
	"gorm.io/gorm"
)

type Connection struct {
	db     *gorm.DB
	driver gorm.Dialector
	config gorm.Config
}

func Make(env environment.Env) (Connection, error) {
	config := gorm.Config{}
	driver := MakeDriver(env)

	db, err := gorm.Open(driver.dialector, &config)

	if err != nil {
		return Connection{}, fmt.Errorf("The database [%s] connection failed: %v\n", env.Get("DATABASE_NAME"), err)
	}

	conn := Connection{
		db:     db,
		config: config,
		driver: driver.dialector,
	}

	return conn, nil
}

func (receiver Connection) GetDB() *gorm.DB {
	return receiver.db
}
