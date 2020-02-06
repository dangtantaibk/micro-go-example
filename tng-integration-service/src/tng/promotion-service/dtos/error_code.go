package dtos

// Common errors code definition.
const (
	UnknownError            = 4000001
	InvalidRequestError     = 4000001
	UnauthorizedError       = 4000002
	ForbiddenError          = 4000003
	InternalServerError     = 4000004
	ServiceUnavailableError = 4000005
)

// Campaign errors
const (
	ErrorUnknown        = 4000000
	ErrorSigNotMatching = 4000401
	ErrorNotFoundData   = 4000402
)
