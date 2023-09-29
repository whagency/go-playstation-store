package ps

import (
	"errors"
)

var (
	ErrResponseFormat = errors.New("wrong response json format in data field")
)
