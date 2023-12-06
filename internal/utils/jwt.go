package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-board/configs"
	"github.com/go-board/internal/dto"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func UserTokenExtract(c *gin.Context) {
	accessMetaData, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("token", accessMetaData)
}

func CreateJwtToken(name string) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": name,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})
	accessSecretKey := os.Getenv("ACCESS_SECRET_KEY")
	accessTokenString, err := accessToken.SignedString([]byte(accessSecretKey))
	if err != nil {
		//tmp log.Error(err)
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	refreshSecretKey := os.Getenv("REFRESH_SECRET_KEY")
	refresgTokenString, err := refreshToken.SignedString([]byte(refreshSecretKey))
	if err != nil {
		//tmp log.Error(err)
		return "", "", err
	}

	redis := configs.ServerConfig.GetRedisClient()
	ctx := context.Background()

	hsetValue := make([]string, 0)
	hsetValue = append(hsetValue, "accessMetaData", accessTokenString, "refreshMetaData", refresgTokenString)
	err = redis.HSet(ctx, name+"_jwtToken", hsetValue).Err()
	if err != nil {
		//tmp log.Error(err)
		return "", "", err
	}
	return accessTokenString, refresgTokenString, nil
}

func JwtTokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString, err := ExtractToken(r)
	if err != nil {
		//tmp log.Error(err)
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//tmp log.Error(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	})
	if err != nil {
		//tmp log.Error(err)
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) (string, error) {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		// if strArr[1] == "null" {
		// 	return "", errors.New("Token format is invalid")
		// }
		return strArr[1], nil
	}
	return "", errors.New("Token format is invalid")
}

func ExtractTokenMetadata(r *http.Request) (*dto.AccessClaims, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		username, ok := claims["username"]
		if !ok {
			return nil, err
		}
		expireTime, ok := claims["exp"]
		if !ok {
			return nil, err
		}
		//tmp log.Infof("username: %s", username.(string))
		//tmp log.Infof("expireTime: %v", expireTime.(float64))

		accessData := dto.AccessClaims{
			Username:   username.(string),
			ExpireTime: expireTime.(float64),
		}
		return &accessData, nil
	}
	return nil, err
}
