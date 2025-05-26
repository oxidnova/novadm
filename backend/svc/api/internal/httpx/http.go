package httpx

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/novadm/backend/driver/schema"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
)

func RespondMessage(c *gin.Context, httpCode int, code code.Code, format string, values ...any) {
	respond(c, httpCode, code, fmt.Sprintf(format, values...), nil)
}

func RespondSuccess(c *gin.Context, d any) {
	respond(c, http.StatusOK, code.Success, "success", d)
}

func RespondItems(c *gin.Context, total int, items any) {
	respond(c, http.StatusOK, code.Success, "success", schema.ArrayItem{Total: total, Items: items})
}

func RespondItemsWithLastId(c *gin.Context, total int, lastId string, items any) {
	respond(c, http.StatusOK, code.Success, "success", schema.ArrayItem{Total: total, Items: items, LastId: lastId})
}

func respond(c *gin.Context, httpCode int, code code.Code, message string, d any) {
	c.JSON(httpCode, schema.Response{
		Code:    int(code),
		Message: message,
		Data:    d,
	})
}
