package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/go-kit/logx"
	"github.com/oxidnova/novadm/backend/driver"
	"github.com/oxidnova/novadm/backend/internal/config"
	"github.com/oxidnova/novadm/backend/svc/api/consultation"
)

// Server is a http RESTFul server for service.
type Server struct {
	engine   *gin.Engine
	d        driver.Registry
	cHandler *consultation.Handler

	httpServer *http.Server
	loading    sync.Mutex
	restarting bool
}

func (s *Server) Load(cfgPath string) error {
	logger := logx.New(logx.DefaultCore, logx.WithName(s.Name()))
	w := &watcher{s: s}

	conf, err := config.LoadWithDefault(cfgPath, w)
	if err != nil {
		return fmt.Errorf("load config, path[%s], err: %v", cfgPath, err)
	}

	s.loading.Lock()
	defer s.loading.Unlock()
	return s.load(logger, &conf)
}

func (s *Server) load(logger logx.Logger, c *config.Config) error {
	d, err := driver.NewRegistry(logger, c)
	if err != nil {
		return err
	}
	s.d = d

	s.setRoutes()
	return nil
}

func (s *Server) Name() string {
	return "nadm.openapi"
}

// Run start server
func (s *Server) Run() {
	s.d.Logger().With("port", s.d.Config().Serve.Api.Port).Info("Starting restful server")

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.d.Config().Serve.Api.Port),
		Handler: s.engine,
	}

	if err := s.httpServer.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) && s.restarting {
			s.d.Logger().Warn("httpd server has closed and is restarting now, please wait...")
		} else {
			s.d.Logger().With("error", err).Fatal("httpd server exceptions")
		}
	}
}

func (s *Server) Stop(sign os.Signal) {
	s.d.Logger().With("signal", sign).Info("Stoping httpd server")
	s.loading.Lock()
	defer s.loading.Unlock()
	s.stop()
}

func (s *Server) stop() {
	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		s.d.Logger().With("error", err).Info("httpd Server forced to shutdown")
	}
	s.d.Logger().Sync()
}

func (s *Server) restart(logger logx.Logger, conf *config.Config) {
	logger.Infof("reloading config & restart service: %v", *conf)

	s.loading.Lock()
	s.restarting = true
	s.stop()
	s.load(logger, conf)
	s.restarting = false
	s.loading.Unlock()

	logger.Infof("success to reloading config & restart service: %v", *conf)
	go s.Run()
}
