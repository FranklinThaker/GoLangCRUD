package models

import(
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Username string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
	jwt.StandardClaims
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}
