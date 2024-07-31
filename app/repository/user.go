package repository

import (
	user_result "fiber-wallet/app/dto/result"
	"fmt"

	"gorm.io/gorm"
)

func InsertUser(tx *gorm.DB, username, password string) int64 {
	var result user_result.InsertUserResult
	fmt.Println(username, "<< username")
	fmt.Println(password, "<< password")
	err := tx.Exec(`INSERT INTO users (username, password) values (?, ?)`, username, password).Select("id").Table("users").Last(&result).Error
	if err != nil {
		panic(err)
	}

	return result.ID
}

// func GetLastInsertedUser(tx *gorm.DB) user_result.GetLastInsertedUserResult {
// 	var result user_result.GetLastInsertedUserResult
// 	err := tx.Raw(`SELECT u.username, LAST_INSERT_ID() AS id FROM users u`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result
// }

func InsertUserWallet(tx *gorm.DB, ID int64) user_result.InsertUserWalletResult {
	var result user_result.InsertUserWalletResult

	err := tx.Raw(`INSERT INTO wallets (user_id) VALUES (?)`, ID).Scan(&result).Error
	if err != nil {
		panic(err)
	}

	return result
}
