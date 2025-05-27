package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/oxidnova/go-kit/logx"
	"github.com/oxidnova/go-kit/sqlxx"
	migrate "github.com/oxidnova/go-kit/sqlxx/migration/sql-migrate"
	"github.com/oxidnova/novadm/backend/internal/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	d dependencies

	db *sqlx.DB
}

type dependencies interface {
	Logger() logx.Logger
	Config() *config.Config
}

func NewStorage(d dependencies) (*Storage, error) {
	dbConf := d.Config().DB
	if err := sqlxx.CreateDBIfNotExist(dbConf.Dsn); err != nil {
		return nil, fmt.Errorf("Couldn't create db, dsn %s, %w", dbConf.Dsn, err)
	}

	// db migrate
	var dbOpts []sqlxx.DBSetupOpt
	if dbConf.MigrationPath != "" {
		dbOpts = append(dbOpts, migrate.Migrate(dbConf.MigrationPath))
	} else {
		dbOpts = append(dbOpts, migrate.MigrateEmbedFS(MigrationsEmbedFS, "/migrations"))
	}

	db, err := sqlxx.Connect(dbConf.Dsn, 3, sqlxx.SetDBSetupOpt(dbOpts...))
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect db, dsn %s, %w", dbConf.Dsn, err)
	}

	db.SetMaxOpenConns(dbConf.MaxOpenConns)
	db.SetMaxIdleConns(dbConf.MaxIdleConns)
	if dbConf.ConnMaxLifeTime > 0 {
		db.SetConnMaxLifetime(dbConf.ConnMaxLifeTime)
	}
	if dbConf.ConnMaxIdleTime > 0 {
		db.SetConnMaxIdleTime(dbConf.ConnMaxIdleTime)
	}

	return &Storage{d: d, db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
