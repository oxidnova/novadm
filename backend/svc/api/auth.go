package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/novadm/backend/driver/schema"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"github.com/oxidnova/novadm/backend/svc/api/internal/httpx"
)

func (s *Server) setAuthRoutes() {
	s.d.Logger().Info("Set Restful Api Router")

	authRouter := s.engine.Group("/api")
	authRouter.POST("/auth/login", s.login)
	authRouter.GET("/user/info", s.userinfo)
	authRouter.POST("/auth/logout", s.logout)
}

func (s *Server) login(c *gin.Context) {
	req := &schema.LoginRequest{}
	if err := c.ShouldBind(&req); err != nil {
		httpx.RespondMessage(c, http.StatusBadRequest, code.InvalidArguments, "invalid paramter for this request.")
		return
	}

	if req.Username == "" {
		httpx.RespondMessage(c, http.StatusBadRequest, code.InvalidArguments, "missing username for this request.")
		return
	}

	if req.Password == "" {
		httpx.RespondMessage(c, http.StatusBadRequest, code.InvalidArguments, "missing password for this request.")
		return
	}

	us, err := s.d.AuthManager().GetUserInfo(req.Username)
	if err != nil {
		httpx.RespondMessage(c, http.StatusBadRequest, code.InvalidArguments, "invalid username or password.")
		return
	}

	if !s.d.AuthManager().CheckCredential(us, strings.TrimSpace(req.Password)) {
		httpx.RespondMessage(c, http.StatusBadRequest, code.InvalidArguments, "invalid username or password.")
		return
	}

	token, err := s.d.AuthManager().ExchangeToken(us)
	if err != nil {
		httpx.RespondMessage(c, http.StatusInternalServerError, code.Internal, "internal server error.")
		return
	}

	if err := s.d.AuthManager().Cookies().IssueToken(c.Writer, token); err != nil {
		httpx.RespondMessage(c, http.StatusInternalServerError, code.Internal, "internal server error.")
		return
	}

	httpx.RespondSuccess(c, schema.LoginToken{AccessToken: "noeffect"})
}

func (s *Server) userinfo(c *gin.Context) {
	// access token cookie
	accessToken, err := s.d.AuthManager().GetTokenFromHttpRequest(c.Request)
	if err != nil {
		httpx.RespondMessage(c, http.StatusInternalServerError, code.Internal, "internal server error.")
		return
	}
	us, err := s.d.AuthManager().VerifyToken(accessToken)
	if err != nil {
		httpx.RespondMessage(c, http.StatusUnauthorized, code.Unauthorized, "Unauthorized")
		return
	}

	httpx.RespondSuccess(c, us)
}

func (s *Server) logout(c *gin.Context) {
	if err := s.d.AuthManager().Cookies().PurgeToken(c.Writer); err != nil {
		httpx.RespondMessage(c, http.StatusInternalServerError, code.Internal, "internal server error.")
		return
	}

	httpx.RespondSuccess(c, nil)
}
