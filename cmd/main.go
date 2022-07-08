package main

import (
	"github.com/diyliv/anime/configs"
	"github.com/diyliv/anime/internal/server"
	"github.com/diyliv/anime/pkg/logger"
	"github.com/diyliv/anime/pkg/storage/postgres"
)

func main() {
	psqlDB := postgres.InitPsqlDB()
	logger := logger.NewLogger(logger.InitLogger())
	configs := configs.ReadConfig()

	server := server.NewServer(*psqlDB, logger, configs)
	server.Run()
}
