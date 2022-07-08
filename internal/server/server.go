package server

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/diyliv/anime/configs"
	searchService "github.com/diyliv/anime/internal/search/delivery/grpc"
	"github.com/diyliv/anime/internal/search/repository"
	"github.com/diyliv/anime/internal/search/usecase"
	"github.com/diyliv/anime/pkg/logger"
	searchpb "github.com/diyliv/anime/proto/animeSearch"
	"google.golang.org/grpc"
)

type server struct {
	psqlDB  sql.DB
	logger  *logger.Logger
	configs *configs.Config
}

func NewServer(psqlDB sql.DB, logger *logger.Logger, configs *configs.Config) *server {
	return &server{psqlDB: psqlDB, logger: logger, configs: configs}
}

func (s *server) Run() {
	s.logger.Info(fmt.Sprintf("Starting gRPC server on port: %s", s.configs.Server.Port))
	lis, err := net.Listen("tcp", s.configs.Server.Port)
	if err != nil {
		s.logger.Error("Error while listening: " + err.Error())
	}

	psqlRepo := repository.NewAnimeRepo(s.psqlDB, s.logger)
	searchUC := usecase.NewSearchUC(psqlRepo)
	searchService := searchService.NewSearchService(s.logger, searchUC)

	opts := []grpc.ServerOption{}

	serv := grpc.NewServer(opts...)
	searchpb.RegisterSearchServiceServer(serv, searchService)

	go func() {
		if err := serv.Serve(lis); err != nil {
			s.logger.Error("Error while serving: " + err.Error())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	serv.GracefulStop()
	s.logger.Info("Exiting was successful")
}
