package dto

import (
	"github.com/pkc918/chat/model"
)

type FriendItemDTO struct {
	model.Account
	Status byte
}
