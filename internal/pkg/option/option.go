package option

import (
	"boilerplate-go/config"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// Option struct for all object that needed
type Option struct {
	Config     config.Provider
	PostgreSQL *sqlx.DB
	Cache      *redis.Client
}
