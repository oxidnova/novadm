package config

import (
	"net/url"

	"github.com/oxidnova/go-kit/configx"
)

var (
	//defaultConfig is default config of the service
	defaultConfig = &Config{
		Version: "1.0.0",
		Serve: Serve{
			Api: ServeRest{
				BaseUrl: &url.URL{},
				Port:    5320,
				Cors: Cors{
					Enabled:          true,
					MaxAge:           12,
					AllowedOrigins:   []string{"*"},
					AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
					AllowedHeaders:   []string{"*", "Authorization"},
					AllowCredentials: true,
				},
			},
		},
	}
)

func LoadWithDefault(configPath string, watcher configx.FileWatcher) (Config, error) {
	var conf Config
	if err := Load(configPath, &conf, defaultConfig, watcher); err != nil {
		return conf, err
	}
	return conf, nil
}

func Load(configPath string, out, defVal interface{}, watcher configx.FileWatcher) error {
	return configx.Load(configPath, out,
		configx.WithDefaultConfig(defVal),
		configx.WithFileWatcher(watcher),
		configx.WithDecodeHookFuncs(
			StringToSameSiteHookFunc(),
		),
	)
}
