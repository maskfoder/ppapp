package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect(connString string) {
	var err error
	Database, err = gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func Migrate() {
	Database.AutoMigrate(&User{}, &Project{}, &Task{})
}

func FindUserByName(name string) (User, error) {
	var user User
	err := Database.Where("Name=?", name).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func FindUserById(id uint) (User, error) {
	var user User
	err := Database.Preload("Projects").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
