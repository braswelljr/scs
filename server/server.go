package server

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

// Server is the http server for the api.
type Server struct {
	Logger   *zerolog.Logger
	Listener net.Listener
	Http     http.Server
}

// Run starts the server that host webapp and api endpoints.
func (s *Server) Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group

	group.Go(func() error {
		<-ctx.Done()
		return s.Http.Shutdown(context.Background())
	})

	group.Go(func() error {
		defer cancel()
		err := s.Http.Serve(s.Listener)
		if err == context.Canceled || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return err
	})

	return group.Wait()
}

// Close closes server and underlying listener.
func (s *Server) Close() error {
	return s.Http.Close()
}
