package service

import (
	"errors"
	"github.com/pkc918/chat/dao"
	"github.com/pkc918/chat/model"
	"github.com/pkc918/chat/response"
	"gorm.io/gorm"
)

func GetFriendList(id string) ([]model.Account, error) {
	friendList, err := dao.GetFriendList(id)
	if err != nil {
		return nil, err
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrRecordNotExits
		}
		return nil, response.ErrInternalServer
	}
	return friendList, nil
}
