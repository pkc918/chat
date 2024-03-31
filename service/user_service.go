package service

import (
	"errors"
	"github.com/pkc918/chat/dao"
	"github.com/pkc918/chat/dto"
	"github.com/pkc918/chat/model"
	"github.com/pkc918/chat/response"
	"gorm.io/gorm"
)

func SignUp(account *dto.SignUpDTO) (*model.Account, error) {

	_account, err := dao.CreateAccount(account)
	if err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, response.ErrUserAlreadyExists
		} else {
			return nil, response.ErrUserRegisterFailed
		}
	}
	return _account, nil
}
