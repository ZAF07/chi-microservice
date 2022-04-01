package domain

import "errors"

type ID string

var (
	ErrEmpty ID = errors.New("empty product ID")
	ErrEmpty Name = errors.New("empty product name")
)

type Product struct {
	id ID 
}