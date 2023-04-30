package helpers

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"time"
)

// JWTManager ...
type JWTManager interface {
	CreateAccessToken(claims *Claims) (string, error)
	ParseJWT(accessToken string) (*Claims, error)
	CreateRefreshToken() string
	IsCorrectJWT(accessToken string) (bool, error)
}

type Claims struct {
	UserId int32 `json:"user_id"`
	jwt.StandardClaims
}

type JWT struct {
	secretKey string
	expiredAt int
}

type RefreshToken struct {
	Token     string
	ExpiredAt int
}

func NewJWT(secretKey string, expiredAt int) (*JWT, error) {
	if secretKey == "" {
		return nil, errors.New("jwt: private key is empty")
	}

	if expiredAt == 0 {
		return nil, errors.New("jwt: expired_at is empty")
	}

	return &JWT{
		secretKey: secretKey,
		expiredAt: expiredAt,
	}, nil
}

func (j *JWT) CreateAccessToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWT) ParseJWT(accessToken string) (*Claims, error) {
	expired := time.Duration(j.expiredAt) * time.Hour
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Add(expired).Unix(),
	}

	token, err := jwt.ParseWithClaims(
		accessToken,
		&Claims{
			StandardClaims: standardClaims,
		},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}

			return []byte(j.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)

	return claims, nil
}

func (j *JWT) IsCorrectJWT(accessToken string) (bool, error) {
	claims, _ := j.ParseJWT(accessToken)

	if claims.ExpiresAt < time.Now().Unix() {
		return false, errors.New("Token lifetime expired")
	}

	token, _ := j.CreateAccessToken(claims)
	if token == accessToken {
		return true, nil
	}

	return false, nil
}

func (j *JWT) CreateRefreshToken() *RefreshToken {
	token := securecookie.GenerateRandomKey(64)
	strToken := base64.StdEncoding.EncodeToString(token)

	return &RefreshToken{
		Token:     strToken,
		ExpiredAt: j.expiredAt,
	}
}
