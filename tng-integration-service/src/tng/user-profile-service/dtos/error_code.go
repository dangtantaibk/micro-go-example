package dtos

// Error define: aaabbb
//		- aaa: module number
//			+ Common: 		100
//			+ UserProfile: 	200
//			+ UserSession: 	300
// 		- bbb: Detail error code

// For Common (100)
const (
	UnknownError               = 100000
	InvalidRequestError        = 100001
	ValidationError            = 100002
	FieldDuplicateError        = 100003
	BusinessLogicError         = 100004
	DateTimeParseError         = 100005
	PermissionInsertError      = 100006
	InvalidBelongRelationError = 100007
	FileIsNotSupportError      = 100008
	FileTooLarge               = 100009
	NotFoundError              = 100010
	MatchingSigError           = 100011
	ConflictError              = 100012 // use for insert duplicate error
	ForbiddenError             = 100013
	InternalServerError        = 100014
)

// For UserProfile (200)
const (
	InsertUserProfileError = 200000
	UpdateUserProfileError = 200001
	GetUserProfileError    = 200002
)

// For UserSession (300)
const (
	LoginInfoInvalid = 300000
	LoginTokenInvalid = 300001
)
