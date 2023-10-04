package database

const (
	GET_BY_DATE = `
		SELECT date, title, url, hd_url, thumb_url, media_type, copyright, explanation, raw_image
			FROM public.images WHERE date = $1 AND deleted_at IS NULL LIMIT 1`

	GET_ALL = `
		SELECT date, title, url, hd_url, thumb_url, media_type, copyright, explanation
			FROM public.images WHERE deleted_at IS NULL`
	
	SAVE_IMAGE = `
	INSERT INTO images
			(date, title, url, hd_url, thumb_url, media_type, copyright, explanation, raw_image) 
 		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`
)
