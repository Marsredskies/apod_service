package envconfig

type Database struct {
	Username string `envconfig:"APOD_PG_USERNAME" default:"postgres"`
	Password string `envconfig:"APOD_PG_PASSWORD" default:"postgres"`
	Host     string `envconfig:"APOD_PG_HOST" default:"postgres"`
	Port     int    `envconfig:"APOD_PG_PORT" default:"5432"`
	DBName   string `envconfig:"APOD_PG_DATABASE_NAME" default:"postgres"`
}
