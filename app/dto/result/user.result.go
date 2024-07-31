package result

type UserRegisterResult struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	WalletID int64  `json:"wallet_id"`
}

type GetLastInsertedUserResult struct {
	ID int64 `json:"id"`
}

type InsertUserResult struct {
	ID int64 `json:"id"`
}

type InsertUserWalletResult struct {
	ID int64 `json:"id"`
}
