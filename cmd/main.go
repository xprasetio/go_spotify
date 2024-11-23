package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_spotify.git/internal/configs"
	"github.com/xprasetio/go_spotify.git/pkg/jwt/internalsql"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	if err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml")); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	cfg = configs.Get()
	_, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	r.Run(cfg.Service.Port)
}
