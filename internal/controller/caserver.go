package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/dto"
	"github.com/go-board/internal/service"
)

type CaServer struct {
	caserver service.CaServer
}

func NewCaServer(caserver service.CaServer) *CaServer {
	return &CaServer{caserver: caserver}
}

func (s *CaServer) CaServerList(c *gin.Context) {
	caservers, err := s.caserver.CaServerList(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, caservers)
}

func (s *CaServer) UserCaServerList(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqUserCaServer)
	caservers, err := s.caserver.UsersCaServerList(c, req.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(200, caservers)
}

func (s *CaServer) CreateCaServer(c *gin.Context) {
	req := c.MustGet("req").(dto.ReqCreateCaServer)
	caserver, err := s.caserver.CreateCaServer(c, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, caserver)
}
