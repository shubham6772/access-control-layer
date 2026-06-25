package jwt

import (
	"errors"

	"access-control-layer/internal/config"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type JwtClient struct {
	secretKey string
}

func NewJwtClient(jwtConf *config.JwtConfig) *JwtClient {
	return &JwtClient{
		secretKey: jwtConf.SecretKey,
	}
}

func (j *JwtClient) Validate(tokenString string) (string, error) {

	token, err := jwtlib.Parse(tokenString, func(token *jwtlib.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwtlib.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id missing")
	}

	return userID, nil
}
