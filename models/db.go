package models

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // mysql driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // mysql driver
	"github.com/johnrazeur/gin-boilerplate/config"
)

// DBInstance is a singleton DB instance
type DBInstance struct {
	Initializer func() *gorm.DB
	instance    interface{}
	once        sync.Once
}

var (
	dbInstance *DBInstance
)

// Instance gets the singleton instance
func (i *DBInstance) Instance() interface{} {
	i.once.Do(func() {
		i.instance = i.Initializer()
	})
	return i.instance
}

func dbInit() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open(config.Config.Database.Dialect, config.Config.Database.DSN)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	// sql log
	if config.Config.Server.Mode == gin.DebugMode {
		db.LogMode(true)
	}

	db.SingularTable(true)

	db.AutoMigrate(&User{})

	return db
}

// DB returns the database instance
func DB() *gorm.DB {
	return dbInstance.Instance().(*gorm.DB)
}

func init() {
	dbInstance = &DBInstance{Initializer: dbInit}
}
