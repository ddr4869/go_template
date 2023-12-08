package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-board/ent"
	"github.com/go-board/internal/controller"
	"github.com/go-board/internal/repository"
	"github.com/go-board/internal/service"
	"github.com/go-board/log"
)

type Server struct {
	Router     *gin.Engine
	controller map[string]interface{}
	middleware map[string]interface{}
}

func NewDefaultServer(clinet *ent.Client) *Server {
	controller := map[string]interface{}{
		"user":     *controller.NewUser(*service.NewUser(*repository.NewUser(clinet))),
		"caserver": *controller.NewCaServer(*service.NewCaServer(*repository.NewCaServer(clinet))),
		"board":    *controller.NewBoard(*service.NewBoard(*repository.NewBoard(clinet))),
		"payment":  *controller.NewPayment(*service.NewPayment(*repository.NewPayment(clinet))),
	}
	middleware := map[string]interface{}{}

	router := NewDefaultGinRouter()

	return &Server{
		Router:     router,
		controller: controller,
		middleware: middleware,
	}
}

func NewDefaultGinRouter() *gin.Engine {
	gin.ForceConsoleColor()

	logStartTime := time.Now()
	logTimeFormat := logStartTime.Format("2006-01-02 15:04:05")
	logPath := os.Getenv("LOG_FILE")
	logFile, err := os.Create(logPath + "-" + logTimeFormat)
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile)

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(GetGinLogFomatter()))

	// cors
	router.Use(corsMiddleware())

	// file max size
	router.MaxMultipartMemory = 8 << 20 // 8MiB

	// redis

	return router
}

func GetGinLogFomatter() gin.LogFormatter {
	return func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor, ginColor string
		var ginColorInt int
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
			ginColorInt, _ = strconv.Atoi(param.StatusCodeColor()[5:7])
			ginColor = fmt.Sprintf("\033[%dm", ginColorInt-10)
		}

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}

		return fmt.Sprintf("%sGIN%s    [%s] |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			ginColor, resetColor,
			param.TimeStamp.Format(time.RFC3339),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			"/Users/ieungyu/go/src/github.com/go-board/logs/log.txt",
			param.ErrorMessage,
		)
	}
}

func corsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}

	config.AddAllowHeaders("Authorization")
	config.AddAllowHeaders("x-frame-options")
	config.AddAllowHeaders("Cache-Control")
	config.AddAllowHeaders("X-XSS-Protection")
	config.AddAllowHeaders("Referrer-Policy")
	config.AddAllowHeaders("Content-Security-Policy")
	config.AddAllowHeaders("Feature-Policy")

	return cors.New(config)
}

func (s *Server) Start(ServerPort string) error {

	srv := &http.Server{
		Addr:    ":" + ServerPort,
		Handler: s.Router,
	}

	log.Infof("Listening and serving HTTP on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("listen: %s\n", err)
		return err
	}

	return nil
}
