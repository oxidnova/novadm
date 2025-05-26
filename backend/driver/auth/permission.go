package auth

import (
	"strings"

	"github.com/oxidnova/novadm/backend/driver/schema"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"github.com/oxidnova/novadm/backend/internal/errorx"
)

var allMenus = map[string]any{
	"all": struct{}{},
	"ses": struct{}{},
	"sms": struct{}{},
	"ac":  struct{}{},
}

func converMenusToFeRoles(menus []string) []string {
	roles := []string{}
	for _, menu := range menus {
		m := strings.TrimSpace(menu)
		if m == "*" {
			m = "all"
		}
		if _, ok := allMenus[m]; ok {
			roles = append(roles, m)
		}
	}

	return roles
}

func (m *defaultManager) CanAccessMenu(us *schema.UserInfo, menu string) error {
	for _, role := range us.Roles {
		if role == "all" {
			return nil
		}

		if role == menu {
			return nil
		}
	}

	return errorx.Errorf(code.Forbidden, "cann't access the menu")
}
