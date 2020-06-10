package model

import "github.com/jinzhu/gorm"

// UserInfo 用户信息
type UserInfo struct {
	gorm.Model
	UserID uint64 `json:"user_id";gorm:"unique;not null"`
	Username string `json:"username";gorm:"unique;not null"`
	Password string	`json:"password"`
	Avatar string  `json:"avatar";gorm:"default:'/avatar/default.png'"`
	BlogID uint64
	Blog Blog
}

func userInfoMigrate()  {
	DB.AutoMigrate(&UserInfo{})
}