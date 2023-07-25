package database

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User db struct and methods
type User struct {
	gorm.Model
	Name     string `gorm:"name;not null;unique" json:"name"`
	Email    string `gorm:"email;not null;unique" json:"email"`
	Password string `gorm:"password" json:"-"`
	Projects []Project
}

func (user *User) Save() (*User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}
	user.Password = string(passwordHash)
	user.Name = html.EscapeString(strings.TrimSpace(user.Name))
	err = Database.Save(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
func (user *User) Delete() (*User, error) {
	err := Database.Delete(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// Project db struct and methods
type Project struct {
	gorm.Model
	ProjName string `gorm:"projname" json:"projname"`
	ProjDesc string `gorm:"projdesc" json:"projdesc"`
	UserID   uint
	Tasks    []Task
}

func (project *Project) Save() (*Project, error) {
	err := Database.Save(&project).Error
	if err != nil {
		return &Project{}, err
	}
	return project, nil
}
func (project *Project) Delete() (*Project, error) {
	err := Database.Delete(&project).Error
	if err != nil {
		return &Project{}, err
	}
	return project, nil
}

// Task db struct and methods
type Task struct {
	gorm.Model
	TaskName   string  `gorm:"taskname" json:"taskname"`
	TaskDesc   string  `gorm:"taskdec" json:"taskdesc"`
	Completion float32 `gorm:"completion" json:"completion"`
	ProjectID  uint
}

func (task *Task) Save() (*Task, error) {
	err := Database.Save(&task).Error
	if err != nil {
		return &Task{}, err
	}
	return task, nil
}
func (task *Task) Delete() (*Task, error) {
	err := Database.Delete(&task).Error
	if err != nil {
		return &Task{}, err
	}
	return task, nil
}
