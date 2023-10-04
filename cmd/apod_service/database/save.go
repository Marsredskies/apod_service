package database

import (
	"context"

	"github.com/Marsredskies/apod_service/cmd/apod_service"
)

func (r *DB) Save(ctx context.Context, i apod.ImageData) error {
	_, err := r.db.ExecContext(ctx, SAVE_IMAGE, i.Date, i.Title, i.URL, i.HDURL, i.ThumbURL, i.MediaType, i.Copyright, i.Explanation, i.RAW)
	return err
}
