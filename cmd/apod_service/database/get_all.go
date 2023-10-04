package database

import (
	"context"
	"database/sql"

	"github.com/Marsredskies/apod_service/cmd/apod_service"
)

func (r *DB) GetAll(ctx context.Context) ([]apod.ImageData, error) {
	var imgs []apod.ImageData

	err := r.db.SelectContext(ctx, &imgs, GET_ALL)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return imgs, err
}
