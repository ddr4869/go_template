package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kafka-go/internal/controller"
	"github.com/kafka-go/internal/middleware"
	// "github.com/the-medium-tech/platform-externals/log"
)

func (s *Server) Routes() {
	r := s.Router
	//api := r.Group("/")
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
	api.GET("/jwt", c.UserJwtTest)
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
	api.GET("/list/user", mwBoard.UserBoardList, c.UserBoardList)
	api.POST("/", mwBoard.CreateBoard, c.CreateBoard)
	api.GET("/recommend", mwBoard.RecommendBoard, c.RecommendBoard)
	api.GET("/hot", c.HotBoardList)
}
