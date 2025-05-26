package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandlerGRPCError(c *gin.Context, err error) {
	errStatus := status.Convert(err)
	if errStatus.Code() == codes.OK {
		return
	}

	switch errStatus.Code() {
	case codes.InvalidArgument, codes.NotFound, codes.AlreadyExists:
		RespondMessage(c, http.StatusBadRequest, code.InvalidArguments, "%s", errStatus.Message())
	case codes.Unauthenticated:
		RespondMessage(c, http.StatusUnauthorized, code.Unauthorized, "Unauthenticated")
	case codes.PermissionDenied:
		RespondMessage(c, http.StatusForbidden, code.Forbidden, "Forbidden")
	default:
		RespondMessage(c, http.StatusInternalServerError, code.Internal, "internal server error.")
	}
}
