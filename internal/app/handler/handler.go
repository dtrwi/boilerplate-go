package handler

import (
	"boilerplate-go/internal/app/service"
	"boilerplate-go/internal/pkg/option"
)

type Option struct {
	option.Option
	*service.Service
}
