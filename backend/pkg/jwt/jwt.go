// pkg/jwt/jwt.go
package jwt

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}

func GenerateToken(userID string, secret string) (string, error) {
    claims := JwtClaims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string) (string, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
        return claims.UserID, nil
    }

    return "", jwt.ErrSignatureInvalid
}