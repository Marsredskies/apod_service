package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/Marsredskies/apod_service/cmd/apod_service"
)

func (r *DB) GetByDate(ctx context.Context, date string) (*apod.ImageData, error) {
	var img apod.ImageData

	err := r.db.GetContext(ctx, &img, GET_BY_DATE, date)
	if err == sql.ErrNoRows {
		log.Println("no rows")
		return nil, nil
	}

	return &img, err
}
