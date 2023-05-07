package middlewares

import (
	"time"

	"github.com/ahsar04/Go-Mini_Project-Ahmad_Saifur_R/constants"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, name string) (string, error) {
	claims :=jwt.MapClaims{}
	claims["user_id"]=userId
	claims["name"]=name
	claims["exp"]=time.Now().Add(time.Hour*1).Unix()

	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
func CreateClientToken(participantId int, name string) (string, error) {
	claims :=jwt.MapClaims{}
	claims["participant_id"]=participantId
	claims["name"]=name
	claims["exp"]=time.Now().Add(time.Hour*1).Unix()

	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.CLIEN_SECRET_JWT))
}
