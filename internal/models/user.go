package models

type TokenData struct {
	Token string `json:"token"`
}

type UserData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
