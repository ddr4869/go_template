package main

import (
	"log"
	"os"

	"github.com/go-board/configs"
	"github.com/go-board/internal/server"
	"github.com/joho/godotenv"
)

// 0 -> .env 파일 등 불러와서 config 세팅
// 1 controller 세팅
// 2 validator 세팅
// 3 repository 세팅
// -> router 세팅
func main() {

	// configs
	_ = godotenv.Load(configs.AppConfigPath)
	cfg := configs.Init()
	entClient := cfg.SetEntRepo()
	_ = cfg.SetRedisRepo()
	cfg.RedisTest()

	server := server.NewDefaultServer(entClient)
	server.Routes()

	if err := server.Start(cfg.Backend.ServerPort); err != nil {
		log.Fatal("staring failed: ", err)
		os.Exit(1)
	}

}
