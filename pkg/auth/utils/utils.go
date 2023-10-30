package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
	"unicode"
)

func GenerateToken(ttl time.Duration, payload interface{}, secretJWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !tok.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims["sub"], nil
}

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func IsValidPassword(password string) bool {
	var (
		upp, low, num bool
		tot           uint8
	)
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		default:
			return false
		}
	}
	if !upp || !low || !num || tot < 8 || tot > 64 {
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}
