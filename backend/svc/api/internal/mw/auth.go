package mw

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oxidnova/go-kit/logx"
	"github.com/oxidnova/novadm/backend/driver/auth"
	"github.com/oxidnova/novadm/backend/driver/schema"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"github.com/oxidnova/novadm/backend/internal/config"
	"github.com/oxidnova/novadm/backend/internal/errorx"
)

type Auth struct {
	d dependencies
}

type dependencies interface {
	Logger() logx.Logger
	Config() *config.Config

	AuthManager() auth.Manager
}

func NewAuth(d dependencies) *Auth {
	return &Auth{d: d}
}

func (m *Auth) HandlerGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		ctx, err := m.VerifyRequest(r)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Request = r.WithContext(ctx)
		c.Next()
	}
}

func (m *Auth) VerifyRequest(r *http.Request) (context.Context, error) {
	// check with token
	ctx, err := m.verifyRequest(r)
	if err == nil {
		return ctx, nil
	}

	return ctx, err
}

var UserInfoCtxKey = struct{}{}

func (m *Auth) verifyRequest(r *http.Request) (context.Context, error) {
	ctx := r.Context()

	// access token cookie
	accessToken, err := m.d.AuthManager().GetTokenFromHttpRequest(r)
	if err != nil {
		return ctx, err
	}

	us, err := m.d.AuthManager().VerifyToken(accessToken)
	if err != nil {
		return ctx, err
	}
	ctx = context.WithValue(ctx, UserInfoCtxKey, us)

	return ctx, nil
}

func UserInfoFromContext(ctx context.Context) *schema.UserInfo {
	ui, _ := ctx.Value(UserInfoCtxKey).(*schema.UserInfo)
	return ui
}

type MonitorAuth struct {
	d dependencies
}

func NewMonitorAuth(d dependencies) *MonitorAuth {
	return &MonitorAuth{d: d}
}

func (m *MonitorAuth) HandlerGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		ctx, err := m.VerifyRequest(r)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Request = r.WithContext(ctx)
		c.Next()
	}
}

func (m *MonitorAuth) VerifyRequest(r *http.Request) (context.Context, error) {
	ctx := r.Context()

	token, err := m.getTokenFromHttpRequest(r)
	if err != nil {
		return ctx, err
	}

	// check with token
	if token != m.d.Config().Serve.Monitor.Token {
		return nil, errorx.Errorf(code.Unauthorized, "token is invalid")
	}

	// return ctx, err
	return ctx, nil
}

func (m *MonitorAuth) getTokenFromHttpRequest(r *http.Request) (string, error) {
	authHeaders, ok := r.Header["Authorization"]
	if !ok {
		return "", errors.New("Authorization header is empty")
	}
	if len(authHeaders) != 1 {
		return "", errors.New("More than one Authorization headers sent")
	}

	return authHeaders[0], nil
}
