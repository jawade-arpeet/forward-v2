package errs

import "errors"

const (
	CodePlatformAlreadyExists = "PLATFORM_ALREADY_EXISTS"
)

var (
	ErrPlatformAlreadyExists = errors.New("platform already exists")
)
