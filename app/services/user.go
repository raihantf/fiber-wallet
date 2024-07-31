package services

import (
	database "fiber-wallet/app/database"
	user_request "fiber-wallet/app/dto/request"
	user_response "fiber-wallet/app/dto/response"
	user_result "fiber-wallet/app/dto/result"
	common_error "fiber-wallet/app/error"
	"fiber-wallet/app/helpers"
	user_repository "fiber-wallet/app/repository"

	"gorm.io/gorm"
)

func UserRegisterService(data *user_request.UserRegisterRequest) (resp *user_response.UserRegisterResponse, err error) {
	var result user_result.UserRegisterResult

	pwd, err := helpers.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	db, err := database.DBInit()
	if err != nil {
		panic(err)
	}

	db.Conn.Transaction(func(tx *gorm.DB) error {
		insertUser := user_repository.InsertUser(tx, data.Username, pwd)
		if insertUser < 1 {
			panic(&common_error.ApiError{
				Message: "Can't insert new user data",
			})
		}

		user_repository.InsertUserWallet(tx, insertUser)

		return nil
	})

	resp = &user_response.UserRegisterResponse{}
	if resp.Token, err = helpers.GenerateToken(&result); err != nil {
		return nil, err
	}

	return resp, err
}
