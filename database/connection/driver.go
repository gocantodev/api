package connection

import (
	"fmt"
	"github.com/voyago/environment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func driver() (gorm.Dialector, *gorm.Config) {
	var datetimePrecision = 2

	return mysql.New(mysql.Config{
		DSN:                       dns(),              //"root:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         256,                // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  false,              // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision, // default datetime precision
		DontSupportRenameIndex:    false,              // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   false,              // use change when rename column, rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,              // smart configure based on used version
	}), &gorm.Config{}
}

func dns() string {
	env, _ := environment.Make("server")

	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", env.Get("DATABASE_HOST"), 3306, "blog")
}
