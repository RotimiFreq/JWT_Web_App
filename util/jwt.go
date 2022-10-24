package util

import (

	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
)


const SECRET_KEY = "hotelbooking"
func GenerateJwt(new_issuer string ) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: new_issuer,
		ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Hour * 24).Unix())),
	})

	token, _ := claims.SignedString(SECRET_KEY)

	return token
	
}

func ParseJWT(current_cookie string ) (string , error) {
	token , err := jwt.ParseWithClaims(current_cookie , &jwt.StandardClaims{}, func(token *jwt.Token)(interface{} , error) {

		return []byte(SECRET_KEY), nil
	})

	if err != nil || !token.Valid  {
		return " ", err 
	}
	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer ,nil 
}


