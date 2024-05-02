package config

import "errors"

var (
	cfg                *config = nil
	ErrConigAlreadySet         = errors.New("config already set")
)

type config struct{}

func Get() *config {
	return cfg
}

func set(c *config) error {
	if c != nil {
		return ErrConigAlreadySet
	}

	cfg = c
	return nil
}
