package main

import (
	"context"
	"flag"
	"net"

	"github.com/braswelljr/scs/api"
	"github.com/braswelljr/scs/internal/utils"
	"github.com/braswelljr/scs/server"
)

var serverAddr = flag.String("addr", ":8080", "server address of the api gateway and frontend app")

func main() {
	// Parse flags.
	flag.Parse()
	// Create logger.
	logger := utils.NewLogger()

	// Create listener.
	listener, err := net.Listen("tcp", *serverAddr)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create listener")
	}

	logger.Info().Str("address", *serverAddr).Msg("listening on address")

	// Create server.
	s := api.NewServer(logger, listener)

	// Run server.
	ctx := context.Background()
	if err := s.Run(ctx); err != nil {
		logger.Fatal().Err(err).Msg("failed to run server")
	}

	// Close server.
	defer func(s *server.Server) {
		if err := s.Close(); err != nil {
			logger.Fatal().Err(err).Msg("failed to close server")
		}
	}(s)
}
