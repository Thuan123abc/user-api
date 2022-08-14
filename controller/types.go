package controller

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
