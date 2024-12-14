package app

import (
	"call-center/internal/delivery/http"
	"call-center/internal/repository"
	"call-center/internal/service"
	"call-center/pkg/database/migration/mymigrate"
	"call-center/pkg/database/postgresql"
	"call-center/pkg/logger/myslog"
	"log/slog"
	"os"
)

func Run() {
	var (
		Port             string = os.Getenv("PORT")
		DatabaseHost     string = os.Getenv("DATABASE_HOST")
		DatabaseUser     string = os.Getenv("DATABASE_USER")
		DatabasePassword string = os.Getenv("DATABASE_PASSWORD")
		DatabaseName     string = os.Getenv("DATABASE_DBNAME")
	)

	mymigrate.Migrate(mymigrate.Params{
		DatabaseHost:     DatabaseHost,
		DatabaseUser:     DatabaseUser,
		DatabasePassword: DatabasePassword,
		DatabaseName:     DatabaseName,
	})
	db, err := postgresql.NewClient(postgresql.Params{
		DatabaseHost:     DatabaseHost,
		DatabaseUser:     DatabaseUser,
		DatabasePassword: DatabasePassword,
		DatabaseName:     DatabaseName,
	})

	if err != nil {
		slog.Error("Failed to connect to database", myslog.ErrAttr(err))
		return
	}

	callRepo := repository.NewCallPostgresqlRepo(db)
	callService := service.NewCallService(callRepo)

	services := service.NewServices(callService)
	handler := http.NewHandler(services)
	handler.Init().Run(":" + Port)
}
