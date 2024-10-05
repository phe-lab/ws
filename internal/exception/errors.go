package exception

import (
	"errors"
)

var ErrNotDirectory = errors.New("workspace path is not directory")
var ErrNotExist = errors.New("workspace path is not exist")
var ErrUserAborted = errors.New("user aborted")
var ErrUnhandled = errors.New("unhandled error")
