package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vpaliy/telex/model"
	"os"
)

type Config struct {
	Type               string
	Path               string
	Name               string
	MaxIdleConnections int
	LogMode            bool
}

func New(config *Config) (*gorm.DB, error) {
	fullPath := config.Path + "/" + config.Name
	db, err := gorm.Open(config.Type, fullPath)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(config.MaxIdleConnections)
	db.LogMode(config.LogMode)
	return db, nil
}

func CreateTestConfig() *Config {
	config := Config{
		Type:               "sqlite3",
		Path:               ".",
		Name:               "test.db",
		MaxIdleConnections: 10,
		LogMode:            true,
	}
	return &config
}

func DropTestDB() error {
	if err := os.Remove("./test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{},
		&model.Subscription{},
		&model.Attachment{},
		&model.Channel{},
		&model.Message{},
		&model.Tag{},
	)
}
