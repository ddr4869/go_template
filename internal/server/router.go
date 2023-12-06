package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-board/internal/controller"
	"github.com/go-board/internal/middleware"
	"github.com/go-board/internal/utils"
	// "github.com/go-board/log"
)

func (s *Server) Routes() {
	r := s.Router
	RouteUser(r, s.controller["user"].(controller.User))
	RouteCaServer(r, s.controller["caserver"].(controller.CaServer))
	RouteBoard(r, s.controller["board"].(controller.Board))
}

func RouteUser(r *gin.Engine, c controller.User) {
	api := r.Group("/user")
	mwUser := middleware.NewUser(c)
	api.GET("/", c.UserList)
	api.POST("/", mwUser.CreateUser, c.CreateUser)
	api.POST("/login", mwUser.UserLogin, c.UserLogin)
	//api.GET("/jwt", c.UserJwtTest)
}

func RouteCaServer(r *gin.Engine, c controller.CaServer) {
	api := r.Group("/caserver")
	mwCaServer := middleware.NewCaServer(c)
	api.GET("/", c.CaServerList)
	api.GET("/user", mwCaServer.UserCaServerList, c.UserCaServerList)
	api.POST("/", mwCaServer.CreateCaServer, c.CreateCaServer)
}

func RouteBoard(r *gin.Engine, c controller.Board) {
	api := r.Group("/board")
	mwBoard := middleware.NewBoard(c)

	api.GET("/view", mwBoard.RecommendBoard, c.ReadBoard)
	api.GET("/list", c.BoardList)
	api.GET("/list/user", utils.UserTokenExtract, c.UserBoardList)
	api.POST("/", mwBoard.CreateBoard, utils.UserTokenExtract, c.CreateBoard)
	api.GET("/recommend", mwBoard.RecommendBoard, c.RecommendBoard)
	api.GET("/hot", c.HotBoardList)
	api.DELETE("/:id", mwBoard.DeleteBoard, utils.UserTokenExtract, c.DeleteBoard)
	api.PATCH("/", mwBoard.PatchBoard, utils.UserTokenExtract, c.PatchBoard)
}
