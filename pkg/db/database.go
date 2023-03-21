// Package init connection to database
package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	gormCfg := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}

	var err error
	var dialect gorm.Dialector

	switch cfg.Driver {
	case "mysql":
		dialect = mysql.Open(cfg.DSN())
	}

	db, err = gorm.Open(dialect, gormCfg)
	if nil != err {
		panic(err)
	}
}

func Open() *gorm.DB {
	if nil != db {
		return db
	}

	return nil
}
