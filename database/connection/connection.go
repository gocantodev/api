package connection

import (
	"fmt"
	"github.com/voyago/environment"
	"gorm.io/gorm"
)

type Connection struct {
	db     *gorm.DB
	driver *gorm.Dialector
}

func Make(env environment.Env) (Connection, error) {
	driver := MakeDriver(env)
	db, err := gorm.Open(driver.dialector, &gorm.Config{})

	if err != nil {
		return Connection{}, fmt.Errorf("The database [%s] connection failed: %v\n", env.Get("DATABASE_NAME"), err)
	}

	conn := Connection{
		db:     db,
		driver: &driver.dialector,
	}

	return conn, nil
}
