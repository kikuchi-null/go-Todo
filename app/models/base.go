package models

import (
	"crypto/sha1"
	"fmt"
	"log"
	"tasks/app/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func gormConnect() *gorm.DB {
	DB, err := gorm.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln()
	}
	return DB
}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plainPassword string) (cryptPassword string) {
	cryptPassword = fmt.Sprintf("%x", sha1.Sum([]byte(plainPassword)))
	return cryptPassword
}

func stringToTime(str string) time.Time {
	layout := "2006-01-02T15:04:05"
	value, _ := time.Parse(layout, str)
	return value
}
