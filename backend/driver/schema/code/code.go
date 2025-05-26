package code

type Code int

const (
	Success          Code = 0 // Success
	BadRequest       Code = 1 // BadRequest
	InvalidArguments Code = 2 // InvalidArguments
	Forbidden        Code = 3 // Forbidden
	NotFound         Code = 4 // NotFound
	Internal         Code = 5 // Internal
	Unauthorized     Code = 6 // Unauthorized
	Unknown          Code = 9 // Unknown
)
