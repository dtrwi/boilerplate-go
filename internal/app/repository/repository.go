package repository

import (
	"boilerplate-go/internal/pkg/option"
)

// Option anything any repository object needed
type Option struct {
	option.Option
}

// Repository all repository object injected here
type Repository struct {
	Cache ICache
}
