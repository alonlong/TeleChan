package db

import (
	"TeleChan/pkg/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// DefaultMaxOpenConns - max open connections for database
	DefaultMaxOpenConns = 20
	// DefaultMaxIdleConns - max idle connections for database
	DefaultMaxIdleConns = 5
)

// Handler for the db operations
type Handler struct {
	settings *config.Config // database connection config
	db       *gorm.DB       // mysql database object
}

// NewHandler returns a new database operation handler
func NewHandler(settings *config.Config) *Handler {
	s := &Handler{settings: settings}

	// init the database connections
	source := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.DB.User,
		settings.DB.Passwd,
		settings.DB.Host,
		settings.DB.Port,
		settings.DB.Name,
	)
	db, err := gorm.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(DefaultMaxOpenConns)
	db.DB().SetMaxIdleConns(DefaultMaxIdleConns)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	// init the db for handler
	s.db = db

	return s
}

// Close make sure all the database connections are released
func (s *Handler) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
