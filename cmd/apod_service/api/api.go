package api

import (
	"context"
	"fmt"

	db "github.com/Marsredskies/apod_service/cmd/apod_service/database"
	"github.com/Marsredskies/apod_service/envconfig"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type API struct {
	echo *echo.Echo
	db   db.Repository
}

func MustInitNewAPI(cnf envconfig.Database) API {
	api, err := New(cnf)
	if err != nil {
		panic(fmt.Errorf("failed to initialize API: %v", err))
	}
	return api
}
func New(cnf envconfig.Database) (API, error) {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	e.Pre(middleware.RemoveTrailingSlash())

	db, err := db.New(cnf)
	if err != nil {
		return API{}, err
	}

	api := API{
		echo: e,
		db:   db,
	}

	e.GET("/album", api.getAllImagesInfo)
	e.GET("/image", api.getImageByDate)

	return api, nil
}

func (a *API) StartServer(port int) error {
	err := a.echo.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return nil
}

func (a *API) Shutdown(ctx context.Context) error {
	return a.echo.Shutdown(ctx)
}
