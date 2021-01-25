package errors

import "errors"

var (
	ErrNotFound = errors.New("port not found")
)

type Error struct {
}
