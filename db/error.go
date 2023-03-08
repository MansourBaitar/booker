package db

import "errors"

var (
	ErrNoConn       = errors.New("no active database connection")
	ErrInvalidQuery = errors.New("invalid query")
	ErrScanFault    = errors.New("unable to scan results")
)
