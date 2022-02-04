package cmd

import (
	"boilerplate-go/config"
	"boilerplate-go/internal/app/appcontext"

	// "boilerplate-go/internal/app/repository"
	// "boilerplate-go/internal/app/server"
	// "boilerplate-go/internal/app/service"
	"boilerplate-go/internal/pkg/option"
)

// Execute adds all child commands.
// This is called by main.main(). It only needs to happen once.
func Execute() {
	var err error
	cfg := config.Config()
	app := appcontext.NewAppContext(cfg)

	var postgreSQL *sqlx.DB
	if cfg.GetBool("postgre.is_enabled") {
		postgreSQL, err = app.GetPostgreSQLConn()
		if err != nil {
			logrus.Fatalf("failed to start, error connect to PostgreSQL | +v", err)
			return
		}
		defer postgreSQL.Close()
	}

	var redis *redis.Client
	if cfg.GetBool("redis.is_enabled") {
		redis, err = app.GetRedisClient()
		if err != nil {
			logrus.Fatalf("failed to start, error connect to Redis | %+v", err)
			return
		}
	}

	option := option.Option{
		Config:     cfg,
		PostgreSQL: postgreSQL,
		Cache:      redis,
	}

	repository := wiringRepository(repository.Option{
		Option: option,
	})

	service := wiringService(service.Option{
		Option:     option,
		Repository: repository,
	})

	server := server.NewServer(option, service)
	server.StartApp()
}

func wiringRepository(option repository.Option) *repository.Repository {
	// wiring up all repositories here
	cacheRepository := repository.NewCache(option)

	repo := repository.Repository{
		Cache: cacheRepository,
	}
	return &repo
}

func wiringService(option service.Option) *service.Service {
	// wiring up all services
	return &svc
}
