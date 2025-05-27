package code

import (
	"github.com/oxidnova/go-kit/x/code"
)

type Code = code.Code

const (
	Success          Code = code.Success          // Success
	BadRequest       Code = code.BadRequest       // BadRequest
	InvalidArguments Code = code.InvalidArguments // InvalidArguments
	Forbidden        Code = code.Forbidden        // Forbidden
	NotFound         Code = code.NotFound         // NotFound
	Internal         Code = code.Internal         // Internal
	Unauthorized     Code = code.Unauthorized     // Unauthorized
	Unknown          Code = code.Unknown          // Unknown
)
