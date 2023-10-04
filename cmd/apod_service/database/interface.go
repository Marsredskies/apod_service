package database

import (
	"context"

	"github.com/Marsredskies/apod_service/cmd/apod_service"
)

type Repository interface {
	GetAll(context.Context) ([]apod.ImageData, error)
	Save(context.Context, apod.ImageData) error
	GetByDate(context.Context, string) (*apod.ImageData, error)
}
