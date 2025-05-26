package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oxidnova/go-kit/ginx/middleware/cors_gin"
	"github.com/oxidnova/go-kit/ginx/middleware/log_gin"
	"github.com/oxidnova/novadm/backend/svc/api/consultation"
)

func (s *Server) setRoutes() {
	engine := gin.New()
	engine.Use(cors_gin.CORS(s.d.Config().Serve.Api.Cors),
		log_gin.Logger(log_gin.Config{
			Logger:              s.d.Logger(),
			ExcluedePrefixPaths: []string{"v2/healthz"},
		}), gin.Recovery())

	engine.GET("/health/ready", checksHealthReady)
	engine.GET("/health/alive", checksHealthAlive)

	s.engine = engine
	s.setAuthRoutes()
	s.ConsultationRoutes().SetRoutes()
	s.setUI()
}

func (s *Server) ConsultationRoutes() *consultation.Handler {
	if s.cHandler == nil {
		s.cHandler = consultation.NewHandler(s.d, s.engine)
	}

	return s.cHandler
}
