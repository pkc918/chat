package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UId      uuid.UUID `gorm:"type:char(36);unique;comment:'id'" json:"UId"`
	Email    string    `gorm:"size:255;not null;unique;comment:'邮箱'" json:"email"`
	Password string    `gorm:"type:text;not null;comment:'密码'" json:"password"`
	Avatar   string    `gorm:"type:text;comment:'头像url'" json:"avatar"`
}

func (act *Account) BeforeCreate(*gorm.DB) (err error) {
	act.UId = uuid.New()
	return
}
