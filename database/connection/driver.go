package connection

import (
	"fmt"
	"github.com/voyago/environment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Driver struct {
	dialector gorm.Dialector
	env       environment.Env
	config    DriverConfig
}

type DriverConfig struct {
	user     string
	password string
	host     string
	port     int
	dbName   string
}

func MakeDriver(env environment.Env) Driver {
	datetimePrecision := 2
	driver := Driver{env: env, config: MakeDriverConfig(env)}

	dialector := mysql.New(mysql.Config{
		DSN:                       driver.GetDns(),    //"root:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  false,              // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    false,              // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   false,              // use change when rename column, rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	})

	driver.dialector = dialector

	return driver
}

func MakeDriverConfig(env environment.Env) DriverConfig {
	port, _ := strconv.Atoi(env.Get("DATABASE_PORT")) //3306

	return DriverConfig{
		user:     env.Get("DATABASE_USER"),
		password: env.Get("DATABASE_PASSWORD"),
		host:     env.Get("DATABASE_HOST"),
		port:     port,
		dbName:   env.Get("DATABASE_NAME"),
	}
}

func (receiver Driver) GetDns() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		receiver.config.user,
		receiver.config.password,
		receiver.config.host,
		receiver.config.port,
		receiver.config.dbName,
	)
}
