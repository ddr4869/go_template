package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/dto"
	"github.com/go-board/internal/repository"
	"github.com/go-board/internal/utils"
	"github.com/the-medium-tech/platform-externals/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	user repository.User
}

func NewUser(user repository.User) *User {
	return &User{user: user}
}

func (s *User) UserList(c *gin.Context) ([]dto.UserDto, error) {
	users, err := s.user.UserList(c)
	if err != nil {
		return nil, err
	}

	dto_users := make([]dto.UserDto, 0)
	for i, v := range users {
		dto_users[i].ID = v.ID
	}

	return dto_users, nil
}

func (s *User) CreateUser(c *gin.Context, name, description, password string) (*dto.User, error) {

	pwByte, err := PasswordEncryption([]byte(password))
	if err != nil {
		return nil, err
	}

	user, err := s.user.CreateUser(c, name, description, pwByte)
	if err != nil {
		return nil, err
	}

	dtoUser := dto.User{
		Name:        user.Name,
		Description: user.Description,
	}

	return &dtoUser, nil
}

func (s *User) UserLogin(c *gin.Context, name, password string) (*dto.User, error) {
	user, err := s.user.GetUser(c, name)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// ACCESS KEY 발급
	accessTokenString, refreshTokenString, err := utils.CreateJwtToken(name)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dto_user := dto.User{
		Name:         user.Name,
		Description:  user.Description,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	return &dto_user, nil
}

func PasswordEncryption(password []byte) (sum []byte, err error) {
	sum, err = bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return sum, nil
}

func CompareHashAndPassword(hashPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashPassword, password)
}
