package configs

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized")
)

type Error struct {
	typeErr error
	appErr  error
}

func NewError(typeErr, appErr error) error {
	return Error{
		typeErr: typeErr,
		appErr:  appErr,
	}
}

func (e Error) Error() string {
	return errors.Join(e.typeErr, e.appErr).Error()
}

func (e Error) TypeError() error {
	return e.typeErr
}

func (e Error) AppError() error {
	return e.appErr
}