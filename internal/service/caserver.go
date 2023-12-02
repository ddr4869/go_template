package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kafka-go/internal/dto"
	"github.com/kafka-go/internal/repository"
)

type CaServer struct {
	caserver repository.CaServer
}

func NewCaServer(caserver repository.CaServer) *CaServer {
	return &CaServer{caserver: caserver}
}

func (s *CaServer) CaServerList(c *gin.Context) ([]dto.CaServerDto, error) {

	caservers, err := s.caserver.CaServerList(c)
	if err != nil {
		return nil, err
	}

	dto_caservers := make([]dto.CaServerDto, 0)
	for _, v := range caservers {
		dto_caserver := dto.CaServerDto{
			ID:   v.ID,
			Name: v.Name,
		}
		dto_caservers = append(dto_caservers, dto_caserver)
	}
	return dto_caservers, nil
}

func (s *CaServer) UsersCaServerList(c *gin.Context, name string) ([]dto.CaServerDto, error) {
	caservers, err := s.caserver.UsersCaServer(c, name)
	if err != nil {
		return nil, err
	}
	dto_caservers := make([]dto.CaServerDto, 0)
	for _, v := range caservers {
		dto_caserver := dto.CaServerDto{
			ID:   v.ID,
			Name: v.Name,
		}
		dto_caservers = append(dto_caservers, dto_caserver)
	}

	return dto_caservers, nil
}

func (s *CaServer) CreateCaServer(c *gin.Context, name string) (*dto.CaServer, error) {
	caserver, err := s.caserver.CreateCaServer(c, name)
	if err != nil {
		return nil, err
	}

	dtoCaServer := dto.CaServer{
		Name: caserver.Name,
	}

	return &dtoCaServer, nil
}
