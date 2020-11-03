package store

import "errors"

var ErrUnknownDriver = errors.New("unknown database driver")

type Store interface {
	Close() error
}
