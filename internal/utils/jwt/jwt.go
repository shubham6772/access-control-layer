package jwt

import (
	"errors"

	"access-control-layer/internal/config"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type JwtClient struct {
	secretKey string
}

type Claims struct {
	UserID    string
	SessionID string
	VideoID   string
}

func NewJwtClient(jwtConf *config.JwtConfig) *JwtClient {
	return &JwtClient{
		secretKey: jwtConf.SecretKey,
	}
}

func (j *JwtClient) Validate(tokenString string) (*Claims, error) {

	token, err := jwtlib.Parse(tokenString, func(token *jwtlib.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	mapClaims, ok := token.Claims.(jwtlib.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	userID, ok := mapClaims["user_id"].(string)
	if !ok {
		return nil, errors.New("user_id missing")
	}

	sessionID, ok := mapClaims["session_id"].(string)
	if !ok {
		return nil, errors.New("session_id missing")
	}

	videoID, ok := mapClaims["video_id"].(string)
	if !ok {
		return nil, errors.New("video_id missing")
	}

	return &Claims{
		UserID:    userID,
		SessionID: sessionID,
		VideoID:   videoID,
	}, nil
}