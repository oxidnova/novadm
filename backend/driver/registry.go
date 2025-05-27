package driver

import (
	"fmt"

	"github.com/oxidnova/go-kit/logx"
	"github.com/oxidnova/novadm/backend/driver/auth"
	"github.com/oxidnova/novadm/backend/internal/config"
	"github.com/oxidnova/novadm/backend/storage"
	"github.com/oxidnova/novadm/backend/storage/sql"
	"go.uber.org/zap"
)

type Registry interface {
	Logger() logx.Logger
	Config() *config.Config

	Storage() storage.Storage
	AuthManager() auth.Manager
}

// NewRegistry retrun a registry
func NewRegistry(logger logx.Logger, c *config.Config) (Registry, error) {
	r := &registry{
		c:      c,
		logger: logger,
	}

	stg, err := sql.NewStorage(r)
	if err != nil {
		return nil, fmt.Errorf("connect to db %w", err)
	}
	r.stg = stg

	logger.With(
		zap.String("endpoint", c.DB.Dsn),
		zap.Int("retryMaxTimes", c.DB.MaxIdleConns),
		zap.Int("retryTimeout", c.DB.MaxOpenConns)).Info("successfully connect to db")

	return r, nil
}

type registry struct {
	c      *config.Config
	logger logx.Logger

	stg         storage.Storage
	authManager auth.Manager
}

func (r *registry) Logger() logx.Logger {
	return r.logger
}

func (r *registry) Config() *config.Config {
	return r.c
}

func (r *registry) Storage() storage.Storage {
	return r.stg
}

func (r *registry) AuthManager() auth.Manager {
	if r.authManager == nil {
		r.authManager = auth.NewManager(r)
	}

	return r.authManager
}
