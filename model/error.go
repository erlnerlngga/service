package model

// Error is for errors in the business domain. See the constants below.
type Error string

const (
	ErrorEmailConflict = Error("EMAIL_CONFLICT")
	ErrorTokenExpired  = Error("TOKEN_EXPIRED")
	ErrorUserInactive  = Error("USER_INACTIVE")
)

func (e Error) Error() string {
	return string(e)
}
