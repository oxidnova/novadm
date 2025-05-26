package auth

import (
	"net/http"

	"github.com/oxidnova/novadm/backend/driver/schema"
)

type CookieManager struct {
	d dependencies

	scheme   string
	sameSite http.SameSite
}

const CookieNameAccessToken = "jwt"

func NewCookieManager(d dependencies) *CookieManager {
	sameSite := d.Config().Auth.Cookies.SameSite
	if sameSite == 0 {
		// default value
		sameSite = http.SameSiteLaxMode
	}

	var scheme string
	if d.Config().Serve.Api.BaseUrl != nil {
		scheme = d.Config().Serve.Api.BaseUrl.Scheme
	}

	return &CookieManager{d: d, scheme: scheme, sameSite: sameSite}
}

func (m *CookieManager) IssueToken(w http.ResponseWriter, token *schema.LoginToken) error {
	m.SetCookie(w, CookieNameAccessToken, token.AccessToken)
	return nil
}

func (m *CookieManager) PurgeToken(w http.ResponseWriter) error {
	m.EmptyCookie(w, CookieNameAccessToken)
	return nil
}

func (m *CookieManager) SetCookie(w http.ResponseWriter, name, value string) {
	var maxAge int
	if m.d.Config().Auth.Token.Lifespan.Seconds() > 0 {
		maxAge = int(m.d.Config().Auth.Token.Lifespan.Seconds())
	}

	m.setCookie(w, name, value, maxAge)
}

func (m *CookieManager) EmptyCookie(w http.ResponseWriter, name string) {
	m.setCookie(w, name, "", -1)
}

func (m *CookieManager) setCookie(w http.ResponseWriter, name, value string, maxAge int) {
	domain := m.d.Config().Auth.Cookies.Domain
	path := m.d.Config().Auth.Cookies.Path
	httpOnly := m.d.Config().Auth.Cookies.HttpOnly

	name, secure := m.secureCookieName(name)
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Domain:   domain,
		Path:     path,
		SameSite: m.sameSite,
		HttpOnly: httpOnly,
		MaxAge:   maxAge,
		Secure:   secure,
	}

	http.SetCookie(w, &cookie)
}

func (m *CookieManager) getCookie(r *http.Request, name string) (string, error) {
	name, _ = m.secureCookieName(name)
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (m *CookieManager) secureCookieName(name string) (string, bool) {
	return secureCookieName(name, m.scheme, m.sameSite)
}

func secureCookieName(name, scheme string, sameSite http.SameSite) (string, bool) {
	if scheme == "https" {
		return "__Secure-" + name, true
	}

	if sameSite == http.SameSiteNoneMode {
		return name, true
	}

	return name, false
}
