package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkc918/chat/dao"
	"github.com/pkc918/chat/dto"
	"github.com/pkc918/chat/model"
	"github.com/pkc918/chat/response"
	"github.com/pkc918/chat/utils"
	"gorm.io/gorm"
	"log"
	"time"
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

func SignIn(account *dto.SignInDTO) (*model.Account, error) {
	_account, err := dao.GetAccount(account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrUserNotFound
		} else {
			return nil, response.ErrInternalServer
		}
	}
	if _account.Password != account.Password {
		return nil, response.ErrPasswordIncorrect
	}
	return _account, nil
}

func CreateJWTToken(account *model.Account) (string, error) {
	userClaims := utils.UserClaims{
		Id:       account.UId,
		Email:    account.Email,
		Password: account.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	signedAccessToken, err := utils.NewAccessToken(userClaims)
	if err != nil {
		log.Fatal("error creating access token")
		return "", err
	}
	return signedAccessToken, nil
}
