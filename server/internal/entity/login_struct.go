package entity

type LoginRequest struct {
	Id       int    `json:"id" gorm:"column:id;primaryKey"`
	Email    string `json:"mail" gorm:"column:mail"`
	Password string `json:"motdepasse" gorm:"column:motdepasse"`
}

func (LoginRequest) TableName() string {
	return "login"
}

type UserResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"mail"`
	Password string `json:"motdepasse"`
}

type LoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
