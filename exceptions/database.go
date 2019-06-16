package exceptions

import (
    "errors"
)

var (
    CONNECTION_ERROR error = errors.New("ConnectionError")
)
