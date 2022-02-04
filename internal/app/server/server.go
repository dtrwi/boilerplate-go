package server

import (
	"boilerplate-go/internal/app/handler"
	"boilerplate-go/internal/app/service"
	"boilerplate-go/internal/pkg/option"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// IServer interface for server
type IServer interface {
	StartApp()
}

type server struct {
	option.Option
	*service.Service
}

// NewServer create object server
func NewServer(opt option.Option, svc *service.Service) IServer {
	return &server{
		Option:  opt,
		Service: svc,
	}
}

func (s *server) StartApp() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderAccept, "x-app-token"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut},
	}))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == s.Config.GetString("app.secret"), nil
		},
	}))

	handler := handler.Option{
		Option:  s.Option,
		Service: s.Service,
	}
	Router(handler, e)

	address := fmt.Sprintf(":%d", s.Config.GetInt("app.port"))
	e.Logger.Fatal(e.Start(address))
}
