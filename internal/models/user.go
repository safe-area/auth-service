package models

type TokenData struct {
	Token string `json:"token"`
}

type AuthResponse struct {
	UserId string `json:"user_id"`
}

type UserData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
