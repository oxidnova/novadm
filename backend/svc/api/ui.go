package api

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) setUI() {
	if s.d.Config().Serve.UI.Dir == "" {
		return
	}

	if err := s.modifyUiConfig(); err != nil {
		s.d.Logger().With("error", err).Fatal("modify UI config: " + s.d.Config().Serve.UI.Dir)
		return
	}

	s.engine.Use(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/api") {
			// 非 /api 请求，交给静态文件服务处理
			if strings.HasSuffix(c.Request.URL.Path, "_app.config.js") {
				// no cache
				c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
				c.Header("Pragma", "no-cache")
				c.Header("Expires", "0")
			} else {
				// cache
				d := s.d.Config().Serve.UI.Lifespan
				c.Header("Cache-Control", fmt.Sprintf("max-age=%d", int(d.Seconds())))
				c.Header("Expires", time.Now().Add(d).Format(http.TimeFormat))
				c.Header("Vary", "Accept-Encoding")
			}

			c.FileFromFS(c.Request.URL.Path, http.Dir(s.d.Config().Serve.UI.Dir))
			c.Abort()
		}
	})
}

func (s *Server) modifyUiConfig() error {
	_, err := os.Stat(s.d.Config().Serve.UI.Dir)
	if err != nil {
		return err
	}

	// use base url as api endpoint
	baseUrl := *s.d.Config().Serve.Api.BaseUrl
	baseUrl.Path = "/api"

	confPath := path.Join(s.d.Config().Serve.UI.Dir, "_app.config.js")
	content := `window._VBEN_ADMIN_PRO_APP_CONF_={"VITE_GLOB_API_URL":"` + baseUrl.String() + `"};` +
		`Object.freeze(window._VBEN_ADMIN_PRO_APP_CONF_);Object.defineProperty(window,"_VBEN_ADMIN_PRO_APP_CONF_",{configurable:false,writable:false,});`
	if err := os.WriteFile(confPath, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}
