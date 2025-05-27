package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oxidnova/go-kit/x/errorx"
	"github.com/oxidnova/novadm/backend/driver/schema"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
)

func (m *defaultManager) ExchangeToken(user *schema.UserInfo) (*schema.LoginToken, error) {
	claim := jwt.NewWithClaims(m.signingMethod, tokenClaims{
		// Payload: p,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.d.Config().Auth.Token.Issuer,
			Subject:   user.Username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.d.Config().Auth.Token.Lifespan)),
		},
	})

	tokenString, err := claim.SignedString([]byte(m.d.Config().Auth.Token.Key))
	if err != nil {
		return nil, err
	}

	return &schema.LoginToken{AccessToken: tokenString}, nil
}

type Payload map[string]any

type tokenClaims struct {
	Payload `json:"payload,omitempty"`
	jwt.RegisteredClaims
}

func (m *defaultManager) VerifyToken(tokenString string) (*schema.UserInfo, error) {
	if tokenString == "" {
		return nil, errorx.Errorf(code.Unauthorized, "token is empty")
	}

	// Parse the signed JWT and validate its contents.
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the key used for signing.
		return []byte(m.d.Config().Auth.Token.Key), nil
	})
	if err != nil {
		return nil, errorx.Errorf(code.Unauthorized, "Failed to parse JWT: %v", err)
	}

	// Check the contents of the parsed JWT.
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errorx.Errorf(code.Unauthorized, "Failed to get claims from token: %v", err)
	}

	us, err := m.GetUserInfo(claims.Subject)
	if err != nil {
		return nil, errorx.Errorf(code.Unauthorized, "User not found")
	}

	return us, nil
}
