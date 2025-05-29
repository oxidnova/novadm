package consultation

import (
	"github.com/gin-gonic/gin"
	"github.com/oxidnova/novadm/backend/driver"
	"github.com/oxidnova/novadm/backend/svc/api/internal/mw"
)

type Handler struct {
	d driver.Registry

	engine *gin.Engine
}

func NewHandler(d driver.Registry, engine *gin.Engine) *Handler {
	return &Handler{d: d, engine: engine}
}

func (h *Handler) SetRoutes() {
	apiRouter := h.engine.Group("/api/consultation")

	authmw := mw.NewAuth(h.d)
	acmw := mw.NewAc(h.d)
	apiRouter.Use(authmw.HandlerGin())
	apiRouter.Use(acmw.HandlerGin())

	// apiRouter.GET("/", h.searchSentEmails)
	// apiRouter.POST("/send", h.sendEmail)
	//
	// // email blacklist
	// apiRouter.GET("/blacklist", h.searchEmailBlacklists)
	// apiRouter.POST("/blacklist", h.addEmailBlacklist)
	// apiRouter.DELETE("/blacklist", h.removeEmailBlacklist)
}
