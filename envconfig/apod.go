package envconfig

type Apod struct {
	BaseURL       string `envconfig:"APOD_NASA_BASE_URL" default:"https://api.nasa.gov/planetary/apod"`
	ApiKey        string `envconfig:"APOD_NASA_API_KEY" default:"DEMO_KEY"`
	IntervalHours int    `envconfig:"APOD_NASA_FETCH_INTERVAL" default:"24"`
	ApiPort       int    `envconfig:"APOD_API_PORT" default:"8080"`
}
