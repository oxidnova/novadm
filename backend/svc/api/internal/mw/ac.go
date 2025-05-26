package mw

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/novadm/backend/driver/schema"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"github.com/oxidnova/novadm/backend/internal/errorx"
)

type Ac struct {
	d dependencies
}

func NewAc(d dependencies) *Ac {
	return &Ac{d: d}
}

func (m *Ac) HandlerGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		us := UserInfoFromContext(r.Context())
		if us == nil {
			c.AbortWithError(http.StatusUnauthorized, errorx.Errorf(code.Unauthorized, "User not found"))
			return
		}

		err := m.VerifyPerimission(r, us)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}
		c.Next()
	}
}

func (m *Ac) VerifyPerimission(r *http.Request, us *schema.UserInfo) error {
	err := m.d.AuthManager().CanAccessMenu(us, getAcMenu(r))
	if err != nil {
		return err
	}

	return nil
}

func getAcMenu(r *http.Request) string {
	path := r.URL.Path
	if idx := strings.Index(path, "/api"); idx != -1 {
		path = path[idx+len("/api"):]
	}

	if strings.HasPrefix(path, "/ses") {
		return "ses"
	}

	if strings.HasPrefix(path, "/sms") {
		return "sms"
	}

	if strings.HasPrefix(path, "/ac") {
		return "ac"
	}

	return ""
}
