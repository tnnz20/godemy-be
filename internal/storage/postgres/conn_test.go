package postgres

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tnnz20/godemy-be/config"
)

func init() {
	err := config.LoadConfig("../../../config/config-local.yaml")
	if err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("success connect to database", func(t *testing.T) {
		db, err := NewConnection(config.Cfg.Database.Postgres)
		if err != nil {
			t.Errorf("error connecting to database: %v", err)
		}

		require.Nil(t, err)
		require.NotNil(t, db)

		defer db.Close()
	})

}
