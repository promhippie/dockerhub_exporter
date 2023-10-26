package dockerhub

// ForbiddenError defines the error if the API endpoint is forbidden.
type ForbiddenError struct{}

// Error implements the errror interface.
func (e ForbiddenError) Error() string {
	return "operation forbidden"
}

// UnknownError defines the error if the API endpoint responds unexpected.
type UnknownError struct{}

// Error implements the errror interface.
func (e UnknownError) Error() string {
	return "authentication issue"
}

// GeneralError defines a generic error handling API responses.
type GeneralError struct {
	message string
}

// Error implements the errror interface.
func (e GeneralError) Error() string {
	return e.message
}
