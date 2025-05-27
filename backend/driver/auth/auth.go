package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oxidnova/go-kit/logx"
	"github.com/oxidnova/go-kit/x/errorx"
	"github.com/oxidnova/novadm/backend/driver/schema"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"github.com/oxidnova/novadm/backend/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type Manager interface {
	GetUserInfo(username string) (*schema.UserInfo, error)
	CheckCredential(user *schema.UserInfo, password string) bool
	ExchangeToken(user *schema.UserInfo) (*schema.LoginToken, error)
	VerifyToken(token string) (*schema.UserInfo, error)
	CanAccessMenu(*schema.UserInfo, string) error
	GetTokenFromHttpRequest(r *http.Request) (string, error)
	Cookies() *CookieManager
}

type defaultManager struct {
	d dependencies

	signingMethod *jwt.SigningMethodHMAC
	cookieManager *CookieManager

	users map[string]*schema.UserInfo
}

func NewManager(d dependencies) Manager {
	var signingMethod *jwt.SigningMethodHMAC
	switch d.Config().Auth.Token.Algorithm {
	case "hs256":
		signingMethod = jwt.SigningMethodHS256
	case "hs512":
		signingMethod = jwt.SigningMethodHS512
	default:
		signingMethod = jwt.SigningMethodHS256
	}

	users := map[string]*schema.UserInfo{}
	for _, cred := range d.Config().Auth.Credentials {
		users[cred.Username] = &schema.UserInfo{
			Username: cred.Username,
			RealName: cred.Realname,
			Password: cred.Password,
			Roles:    converMenusToFeRoles(cred.Menus),
		}
	}

	return &defaultManager{d: d, signingMethod: signingMethod, users: users}
}

type dependencies interface {
	Logger() logx.Logger
	Config() *config.Config
}

func (m *defaultManager) GetUserInfo(username string) (*schema.UserInfo, error) {
	user, ok := m.users[strings.TrimSpace(username)]
	if !ok {
		return nil, errorx.Errorf(code.NotFound, "user not found")
	}

	return user, nil
}

func (m *defaultManager) CheckCredential(user *schema.UserInfo, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (m *defaultManager) GenerateHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (m *defaultManager) GetTokenFromHttpRequest(r *http.Request) (string, error) {
	// access token from cookie
	accessToken, err := m.Cookies().getCookie(r, CookieNameAccessToken)
	if err == nil && accessToken != "" {
		return accessToken, nil
	}

	// access token from authorization header
	accessToken, err = extractTokenFromHeader(r)
	if err == nil && accessToken != "" {
		return accessToken, nil
	}

	return "", nil
}

func extractTokenFromHeader(r *http.Request) (string, error) {
	authHeaders, ok := r.Header["Authorization"]
	if !ok {
		return "", errors.New("Authorization header is empty")
	}
	if len(authHeaders) != 1 {
		return "", errors.New("More than one Authorization headers sent")
	}

	parts := strings.SplitN(authHeaders[0], " ", 2)
	if len(parts) != 2 {
		return "", errors.New("Bad Authorization header")
	}
	if !strings.EqualFold(parts[0], "bearer") {
		return "", errors.New("Only Bearer tokens accepted")
	}
	return parts[1], nil
}

func (m *defaultManager) Cookies() *CookieManager {
	if m.cookieManager == nil {
		m.cookieManager = NewCookieManager(m.d)
	}

	return m.cookieManager
}
