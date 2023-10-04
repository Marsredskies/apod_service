package nasa

import (
	"context"
	"testing"

	"github.com/Marsredskies/apod_service/cmd/apod_service/database"
	"github.com/Marsredskies/apod_service/envconfig"
	"github.com/stretchr/testify/require"
)

func TestMigration(t *testing.T) {
	ctx := context.Background()

	cnf := envconfig.Database{
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     5432,
		DBName:   "postgres",
	}
	require.NoError(t, database.ApplyMigrations(ctx, cnf))
}
