package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kafka-go/internal/dto"
	"github.com/kafka-go/internal/repository"
)

type Admin struct {
	admin repository.Admin
}

func NewAdmin(admin repository.Admin) *Admin {
	return &Admin{admin: admin}
}

func (s *Admin) AdminList(c *gin.Context) ([]dto.AdminDto, error) {
	Admins, err := s.admin.AdminList(c)
	if err != nil {
		return nil, err
	}

	dto_Admins := make([]dto.AdminDto, 0)
	for i, v := range Admins {
		dto_Admins[i].ID = v.ID
	}

	return dto_Admins, nil
}

func (s *Admin) CreateAdmin(c *gin.Context, name, description string) (*dto.Admin, error) {
	Admin, err := s.admin.CreateAdmin(c, name)
	if err != nil {
		return nil, err
	}

	dtoAdmin := dto.Admin{
		Name: Admin.Name,
		//Description: Admin.Description,
	}

	return &dtoAdmin, nil
}
