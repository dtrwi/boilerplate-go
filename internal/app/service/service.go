package service

import (
	"boilerplate-go/internal/app/repository"
	"boilerplate-go/internal/pkg/option"
)

// Option anything any service object needed
type Option struct {
	option.Option
	*repository.Repository
}

// Service all service object injected here
type Service struct {
}
