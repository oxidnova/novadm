package driver

import (
	"github.com/oxidnova/go-kit/logx"
	"github.com/oxidnova/novadm/backend/driver/auth"
	"github.com/oxidnova/novadm/backend/internal/config"
)

type Registry interface {
	Logger() logx.Logger
	Config() *config.Config

	AuthManager() auth.Manager
}

// NewRegistry retrun a registry
func NewRegistry(logger logx.Logger, c *config.Config) (Registry, error) {

	return &registry{
		c:      c,
		logger: logger,
	}, nil
}

type registry struct {
	c      *config.Config
	logger logx.Logger

	authManager auth.Manager
}

func (r *registry) Logger() logx.Logger {
	return r.logger
}

func (r *registry) Config() *config.Config {
	return r.c
}

func (r *registry) AuthManager() auth.Manager {
	if r.authManager == nil {
		r.authManager = auth.NewManager(r)
	}

	return r.authManager
}
