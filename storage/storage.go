package storage

import "errors"

var (
	ErrClientNotFound = errors.New("client not found")
	ErrClientExists   = errors.New("cleint exists")
)
