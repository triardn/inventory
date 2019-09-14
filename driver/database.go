package driver

import (
	"time"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBOption struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	DBName               string
	AdditionalParameters string
	ConnectionSetting    ConnectionSetting
}

type ConnectionSetting struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewSqliteDatabase(pathToDB string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", pathToDB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
