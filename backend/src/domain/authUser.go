package domain

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// パスワードハッシュ化
func PasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error: ", err)
		return "", fmt.Errorf(InternalServerError)
	}
	fmt.Println("\n" + string(hashedPassword) + "\n")
	return string(hashedPassword), nil
}

// パスワード認証
func ComparePassword(requestPassword string, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword)); err != nil {
		log.Println("error: ", err)
		return fmt.Errorf(Unauthorized)
	}
	return nil
}

type Claims struct {
	UserId string `json:"userid"`
	jwt.RegisteredClaims
}

// トークン発行
func CreateToken(userId int) (string, error) {
	time := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserId: strconv.Itoa(userId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time),
		},
	}

	// トークン作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

// トークンの検証
func ValidateToken(requestToken string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(requestToken, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_KEY"), nil
	})
	if err != nil || !token.Valid {
		log.Println("error: ", err)
		return nil, fmt.Errorf(Unauthorized)
	}

	return claims, nil
}
