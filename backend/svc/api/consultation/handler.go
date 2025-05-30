package consultation

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/go-kit/x/errorx"
	"github.com/oxidnova/novadm/backend/driver"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"github.com/oxidnova/novadm/backend/storage"
	"github.com/oxidnova/novadm/backend/svc/api/internal/httpx"
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

	apiRouter.GET("/", h.searchConsultations)
	apiRouter.POST("/crawl", h.crawlConsultation)
}

type searchConsultationsParams struct {
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
	ID        string `json:"id" form:"id"`
	Status    int    `json:"status" form:"status"`
	StartTime int64  `json:"startTime" form:"startTime"`
	EndTime   int64  `json:"endTime" form:"endTime"`
}

func (h *Handler) searchConsultations(c *gin.Context) {
	var params searchConsultationsParams
	if err := c.ShouldBindQuery(&params); err != nil {
		httpx.RespondMessage(c, http.StatusBadRequest, code.InvalidArguments, "invalid arguments for this request.")
		return
	}

	var (
		total         int
		consultations []*storage.CrossConsultation
	)
	if params.ID != "" {
		consultation, err := h.d.Storage().GetCrossConsultationByID(params.ID)
		if err != nil {
			errStatus := errorx.ConvertError(err)
			if errStatus.Code == code.NotFound {
				httpx.RespondItems(c, total, consultations)
				return
			}

			h.d.Logger().Info("get consultation: " + err.Error())
			httpx.HandlerError(c, err)
			return
		}

		total = 1
		consultations = append(consultations, consultation)
	} else {
		var err error
		consultations, total, err = h.d.Storage().ListCrossConsultationsByFilter(
			toListCrossConsultationsByFilter(params),
		)
		if err != nil {
			h.d.Logger().Info("search consultations: " + err.Error())
			httpx.HandlerError(c, err)
			return
		}
	}

	httpx.RespondItems(c, total, consultations)
}

func toListCrossConsultationsByFilter(params searchConsultationsParams) *storage.CrossConsultationFilter {
	where := &storage.CrossConsultationFilter{
		Offset: (params.Page - 1) * params.PageSize,
		Limit:  params.PageSize,
		Status: params.Status,
	}
	if params.StartTime > 0 {
		where.StartTime = time.Unix(params.StartTime, 0)
	}
	if params.EndTime > 0 {
		where.EndTime = time.Unix(params.EndTime, 0)
	}
	if where.Offset < 0 {
		where.Offset = 0
	}
	if where.Limit <= 0 {
		where.Limit = 50
	}

	return where
}

func (h *Handler) crawlConsultation(c *gin.Context) {
}
