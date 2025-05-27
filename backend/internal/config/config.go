package config

import (
	"net/http"
	"net/url"
	"time"

	"github.com/oxidnova/go-kit/ginx/middleware/cors_gin"
)

// Conf is the entry of the configuration file
type Config struct {
	Version string `koanf:"version"`
	Serve   Serve  `koanf:"serve"`
	Auth    Auth   `koanf:"auth"`
	DB      DB     `koanf:"db"`

	Development bool `koanf:"dev"`
}

// Serve for service configuration
type Serve struct {
	Api ServeRest `koanf:"api"`
	UI  ServeUI   `koanf:"ui"`
}

// ServeRest for Restful configuration
type ServeRest struct {
	BaseUrl *url.URL `koanf:"base_url"`
	Port    int      `koanf:"port"`
	Cors    Cors     `koanf:"cors"`
}

type ServeUI struct {
	Dir      string        `koanf:"dir"`
	Lifespan time.Duration `koanf:"lifespan"`
}

type Auth struct {
	Token       Token        `koanf:"token"`
	Cookies     Cookies      `koanf:"cookies"`
	Credentials []Credential `koanf:"credentials"`
}

type Token struct {
	Issuer    string        `koanf:"issuer"`
	Lifespan  time.Duration `koanf:"lifespan"`
	Algorithm string        `koanf:"algorithm"`
	Key       string        `koanf:"key"`
}

type Cookies struct {
	Path     string        `koanf:"path"`
	SameSite http.SameSite `koanf:"same_site"`
	Domain   string        `koanf:"domain"`
	Lifespan time.Duration `koanf:"lifespan"`
	HttpOnly bool          `koanf:"http_only"`
}

type Credential struct {
	Username string   `koanf:"username"`
	Realname string   `koanf:"realname"`
	Password string   `koanf:"password"`
	Menus    []string `koanf:"menus"`
}

// Cors for cross domain config
type Cors = cors_gin.Config

type DB struct {
	Dsn             string        `koanf:"dsn"`
	MigrationPath   string        `koanf:"migration_path"`
	MaxIdleConns    int           `koanf:"max_idle_conns"`
	MaxOpenConns    int           `koanf:"max_open_conns"`
	ConnMaxLifeTime time.Duration `koanf:"conn_max_life_time"`
	ConnMaxIdleTime time.Duration `koanf:"conn_max_idle_time"`
}
