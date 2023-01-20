package user

import "github.com/google/wire"

const (
	UID              string          = "uid"
	PasswordNotMatch CredentialError = CredentialError("PasswordNotMatch")
)

var HandlerSet = wire.NewSet(
	AuthHandlerSet,
	InfoHandlerSet,
)

type CredentialError string

func (e CredentialError) Error() string {
	return string("UserCredentialError: " + e)
}

func (CredentialError) UserCredentialError() {}
