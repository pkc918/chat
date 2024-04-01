package dao

import (
	"github.com/google/uuid"
	"github.com/pkc918/chat/db"
	"github.com/pkc918/chat/dto"
	"github.com/pkc918/chat/model"
)

func CreateAccount(dto *dto.SignUpDTO) (*model.Account, error) {
	account := &model.Account{
		UId:      uuid.New(),
		Email:    dto.Email,
		Password: dto.Password,
		Avatar:   dto.Avatar,
	}

	err := db.DB.Model(&model.Account{}).Create(account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func GetAccount(dto *dto.SignInDTO) (*model.Account, error) {
	account := &model.Account{}
	err := db.DB.Model(&model.Account{}).Where("email = ?", dto.Email).First(account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}
