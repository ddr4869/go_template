package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/dto"
	"github.com/go-board/internal/service"
)

type User struct {
	user service.User
}

func NewUser(user service.User) *User {
	return &User{user: user}
}

// Controller
func (s *User) UserList(c *gin.Context) {
	users, err := s.user.UserList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, users)
}

func (s *User) UserLogin(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqLoginUser)
	user, err := s.user.UserLogin(c, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

func (s *User) CreateUser(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqCreateUser)
	user, err := s.user.CreateUser(c, req.Name, req.Description, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
