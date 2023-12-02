package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kafka-go/configs"
	"github.com/the-medium-tech/platform-externals/log"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func CreateJwtToken(name string) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": name,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	})
	accessSecretKey := os.Getenv("ACCESS_SECRET_KEY")
	accessTokenString, err := accessToken.SignedString([]byte(accessSecretKey))
	if err != nil {
		log.Error(err)
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	refreshSecretKey := os.Getenv("REFRESH_SECRET_KEY")
	refresgTokenString, err := refreshToken.SignedString([]byte(refreshSecretKey))
	if err != nil {
		log.Error(err)
		return "", "", err
	}

	redis := configs.ServerConfig.GetRedisClient()
	ctx := context.Background()

	hsetValue := make([]string, 0)
	hsetValue = append(hsetValue, "accessMetaData", accessTokenString, "refreshMetaData", refresgTokenString)
	err = redis.HSet(ctx, name+"_jwtToken", hsetValue).Err()
	if err != nil {
		log.Error(err)
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
	tokenString := ExtractToken(r)
	log.Info("2")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	})
	log.Info("3")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	log.Infof("bearToken: ", bearToken)
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	log.Info("1")
	return ""
}

func ExtractTokenMetadata(r *http.Request) ([]string, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	log.Infof("token: %+v", token)
	claims, ok := token.Claims.(jwt.MapClaims)
	log.Infof("claims: %+v", claims)
	if ok && token.Valid {
		username, ok := claims["username"]
		if !ok {
			return nil, err
		}
		expireTime, ok := claims["exp"]
		if !ok {
			return nil, err
		}
		log.Infof("username: %s", username)
		log.Infof("expireTime: %v", expireTime)

		return nil, nil
	}
	return nil, err
}
