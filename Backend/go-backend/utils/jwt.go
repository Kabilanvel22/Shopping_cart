package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)


var jwtSecret = []byte("n-Wvoe3gxZzzfsrw9dTsh93slcBRoI4g89P6LQzzl24") 


func ParseToken(tokenString string) (*jwt.Token, error) {
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
