package model

import (
	"blog-service/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewEngine(d *setting.DataBaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", d.UserName, d.PassWord, d.Host, d.DBName, d.Charset, d.ParseTime)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		Logger: newLogger,
	})

	sqlDb, _ := db.DB()

	err = sqlDb.Ping()

	if err != nil {
		return nil, err
	}
	return db, nil
}
