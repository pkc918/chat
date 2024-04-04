package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FriendShip struct {
	gorm.Model
	UserId   uuid.UUID `gorm:"type:char(36);comment:'user id'" json:"userId"`
	FriendId uuid.UUID `gorm:"type:char(36);comment:'friend id'" json:"friendId"`
	Status   byte      `gorm:"size:1;comment:'1:待确认,2:已确认,3:已删除'" json:"status"`
}
