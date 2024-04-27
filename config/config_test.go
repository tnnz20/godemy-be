package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("success load config", func(t *testing.T) {
		err := Load("config-local.yaml")

		require.Nil(t, err)
		log.Printf("config: %+v", Cfg)
	})

	t.Run("failed load config", func(t *testing.T) {
		err := Load("config-local-err.yaml")

		require.NotNil(t, err)
		log.Printf("error: %+v", err)
	})
}
