package utils

import (
    "github.com/dgrijalva/jwt-go"
    "os"
    "time"
)

func GenerateToken(userId uint) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": userId,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenString string) (uint, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return uint(claims["userId"].(float64)), nil
    }
    return 0, err
}
