package models

import (
	"crypto/sha1"
	"fmt"
	"log"
	"tasks/app/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var Driver string = config.Config.SQLDriver
var DBname string = config.Config.DBName

func gormConnect() *gorm.DB {
	DB, err := gorm.Open(Driver, DBname)
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
