package response

const (
	StatusInternalServerName  = "internal_server_error"
	InternalServerDescription = "The server is unable to complete your request"
)

const (
	StatusUnauthorizedErrorName        = "access_denied"
	StatusUnauthorizedErrorDescription = "Authorization failed: token doesn't exist or invalid"
	StatusUnauthorizedInvalidToken     = "Invalid token"
)

const (
	StatusForbidden            = "forbidden"
	StatusForbiddenDescription = "You are not allowed to access this API"
)

const (
	StatusNotFound            = "not_found"
	StatusNotFoundDescription = "The requested resource is not found"
)

const (
	StatusConflict            = "conflict"
	StatusConflictDescription = "The request could not be completed due to a conflict with the current state of the target resource"
)

const (
	StatusBadRequestErrorName        = "bad_request"
	StatusBadRequestErrorDescription = "Your request resulted in error"
)

const (
	StatusSuccessCreatedName = "success_created"
	StatusSuccessOK          = "success_ok"
	StatusSuccessLogin       = "success_login"
)
