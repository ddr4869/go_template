package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/controller"
	"github.com/go-board/internal/dto"
)

type User struct {
	user controller.User
}

func NewUser(user controller.User) *User {
	return &User{user: user}
}

func (s *User) CreateUser(c *gin.Context) {
	var req dto.ReqCreateUser
	if err := c.ShouldBind(&req); err != nil {
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *User) UserLogin(c *gin.Context) {
	var req dto.ReqLoginUser
	if err := c.ShouldBind(&req); err != nil {
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}
