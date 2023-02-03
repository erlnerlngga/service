package model

// Error is for errors in the business domain. See the constants below.
type Error string

const (
	ErrorEmailConflict = Error("EMAIL_CONFLICT")
)

func (e Error) Error() string {
	return string(e)
}
