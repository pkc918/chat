package dto

import (
	"github.com/google/uuid"
	"time"
)

type FriendItemDTO struct {
	CreatedAt time.Time
	Uid       uuid.UUID
	Email     string
	Avatar    string
	Status    byte
}
