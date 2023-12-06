package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/controller"
	"github.com/go-board/internal/dto"
)

type CaServer struct {
	caserver controller.CaServer
}

func NewCaServer(caserver controller.CaServer) *CaServer {
	return &CaServer{caserver: caserver}
}

func (s *CaServer) CreateCaServer(c *gin.Context) {
	var req dto.ReqCreateCaServer
	if err := c.ShouldBind(&req); err != nil {
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *CaServer) UserCaServerList(c *gin.Context) {
	var req dto.ReqUserCaServer
	if err := c.ShouldBind(&req); err != nil {
		//tmp log.Info(err)
		MwAbortJson(c, http.StatusBadRequest, err)
		return
	}
	c.Set("req", req)
	c.Next()
}
