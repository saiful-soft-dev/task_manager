package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("mysecretkey")

func GenerateJWT(userID uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (uint, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil || !token.Valid {
        return 0, err
    }

    claims := token.Claims.(jwt.MapClaims)
    return uint(claims["user_id"].(float64)), nil
}
