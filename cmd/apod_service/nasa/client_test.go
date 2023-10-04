package nasa

import (
	"context"
	"testing"
	"time"

	"github.com/Marsredskies/apod_service/cmd/apod_service/database"
	"github.com/Marsredskies/apod_service/envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSaveImage(t *testing.T) {
	ctx := context.Background()

	dbConfig := envconfig.Database{
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     5432,
		DBName:   "postgres",
	}

	clientConfig := envconfig.Apod{
		BaseURL:       "https://api.nasa.gov/planetary/apod",
		ApiKey:        "DEMO_KEY",
		IntervalHours: 24,
	}

	db, err := database.ConnectDB(ctx, dbConfig)
	require.NoError(t, err)

	database.DropMigrations(db)
	database.MustApplyMigrations(ctx, dbConfig)

	dbClient, err := database.New(dbConfig)
	require.NoError(t, err)

	n, err := InitClient(clientConfig, dbClient)
	require.NoError(t, err)

	require.NoError(t, n.FetchAndSaveAPOD(ctx, time.Now().AddDate(0, 0, -1)))

	img, err := n.db.GetByDate(ctx, time.Now().AddDate(0, 0, -1).Format(dateFormat))
	require.NoError(t, err)
	require.NotNil(t, img)

	imgs, err := n.db.GetAll(ctx)
	require.NoError(t, err)
	require.NotEqual(t, 0, len(imgs))
}

func TestGetExtension(t *testing.T) {
	n := NasaClient{}
	testCases := []struct {
		url         string
		expectedExt string
		expectedErr bool
	}{
		{
			url:         "https://test.domain.com/library/image.jpg?width=400&height=700",
			expectedExt: "jpg",
			expectedErr: false,
		}, {
			url:         "https://example.com/nothing",
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			ext, err := n.GetFileExtensionFromUrl(tc.url)
			if tc.expectedErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, ext, tc.expectedExt)
			}
		})
	}
}
