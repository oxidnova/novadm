package config

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/oxidnova/novadm/backend/internal/x"
)

// StringToSameSiteHookFunc returns a DecodeHookFunc that converts
// strings to http.SameSite.
func StringToSameSiteHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(http.SameSiteStrictMode) {
			return data, nil
		}

		// Convert it by parsing
		s, ok := data.(string)
		if !ok {
			return data, nil
		}

		if s == "" {
			s = "lax"
		}

		var mode http.SameSite
		switch cs := x.SwitchExact(s); {
		case cs.AddCase("strict"):
			mode = http.SameSiteStrictMode
		case cs.AddCase("lax"):
			mode = http.SameSiteLaxMode
		case cs.AddCase("none"):
			mode = http.SameSiteNoneMode
		default:
			return nil, fmt.Errorf("Unable to same site: %s", cs.ToUnknownCase())
		}

		return mode, nil
	}
}
