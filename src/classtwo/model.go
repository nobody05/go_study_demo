package classtwo


import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
 * @Author: gaoz
 * @Date: 2020/9/28
 */

type GoUser struct {
	gorm.Model
	Id int `gorm:"primary_key"`
	Username string `json:"username";gorm:"column:username"`
	Password string `json:"password";gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status int8
}

func (u GoUser)TableName() string {
	return "go_user"
}

func (u GoUser) CreateUser(db *gorm.DB, user GoUser) int {
	db.Create(user)
	return user.Id
}