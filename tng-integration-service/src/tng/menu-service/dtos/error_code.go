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

// User errors
const (
	ErrorNotFound = 5000404
	ErrorPasswordIncorrect = 5000401
	ErrorPhoneNumberNotFound = 5000402
)

// H5ZaloPay errors
const (
	ErrorUnknown           = 4000000
	ErrorSigNotMatching    = 4000401
	ErrorNotFoundData      = 4000402
	ErrorMBTokenIsChange   = 4000403
	ErrorDataOfH5Incorrect = 4000404
	ErrorMBTokenIsBlock    = 4000405
)
