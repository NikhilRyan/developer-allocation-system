package utils

import (
    "errors"
)

var (
    ErrNotFound     = errors.New("resource not found")
    ErrUnauthorized = errors.New("unauthorized access")
    ErrBadRequest   = errors.New("bad request")
)
