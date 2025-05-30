package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/go-kit/x/errorx"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
)

func HandlerError(c *gin.Context, err error) {
	errStatus := errorx.ConvertError(err)
	if errStatus.Code == code.Success {
		return
	}

	switch errStatus.Code {
	case code.InvalidArguments, code.NotFound:
		RespondMessage(c, http.StatusBadRequest, code.InvalidArguments, "%s", err.Error())
	case code.Unauthorized:
		RespondMessage(c, http.StatusUnauthorized, code.Unauthorized, "Unauthenticated")
	case code.Forbidden:
		RespondMessage(c, http.StatusForbidden, code.Forbidden, "Forbidden")
	default:
		RespondMessage(c, http.StatusInternalServerError, code.Internal, "internal server error.")
	}
}
