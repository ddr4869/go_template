package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kafka-go/internal/dto"
	"github.com/kafka-go/internal/service"
)

type Admin struct {
	admin service.Admin
}

func NewAdmin(admin service.Admin) *Admin {
	return &Admin{admin: admin}
}

// Controller
func (s *Admin) AdminList(c *gin.Context) {
	Admins, err := s.admin.AdminList(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, Admins)
}

func (s *Admin) CreateAdmin(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqCreateAdmin)
	Admin, err := s.admin.CreateAdmin(c, req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, Admin)
}
