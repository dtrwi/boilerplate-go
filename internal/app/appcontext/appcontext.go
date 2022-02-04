package appcontext

import (
	"boilerplate-go/config"
	"boilerplate-go/internal/app/driver"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// AppContext the app context struct
type AppContext struct {
	config config.Provider
}

// NewAppContext initiate appcontext object
func NewAppContext(config config.Provider) *AppContext {
	return &AppContext{
		config: config,
	}
}

func (a *AppContext) GetPostgreSQLConn() (*sqlx.DB, error) {
	return driver.NewPostgreSQL(
		driver.PostgreSQLOption{
			URL:         a.config.GetString("postgre.url"),
			MaxIdleConn: a.config.GetInt("postgre.max_idle_connections"),
			MaxOpenConn: a.config.GetInt("postgre.max_open_connections"),
		})
}

func (a *AppContext) GetRedisClient() (redisClient *redis.Client, err error) {
	redisClient, err = driver.NewRedis(
		driver.RedisOption{
			URL: a.config.GetString("redis.url"),
		},
	)
	return
}
