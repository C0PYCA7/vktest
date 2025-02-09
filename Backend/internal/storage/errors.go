package storage

import "errors"

var (
	ErrUniqueIP = errors.New("container with that ip already exists")
	ErrBeginTx  = errors.New("begin transaction error")
)
