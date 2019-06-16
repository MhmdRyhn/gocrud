package exceptions

import (
    "errors"
)

var (
    KEY_ERROR error = errors.New("KeyError")
    TYPE_ERROR error = errors.New("TypeError")
)
