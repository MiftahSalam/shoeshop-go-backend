package database_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"shoeshop-backend/src/infrastructure/logger"
	"shoeshop-backend/src/shared/config"
	"shoeshop-backend/src/shared/database"
)

func TestPostgresConnect(t *testing.T) {
	asserts := assert.New(t)

	os.Setenv("CONFIG_FILE", "../../../.env")
	appConfig := config.Setup()
	log := logger.NewLogger(&appConfig.Logger)
	db := database.Setup(appConfig.Database, &log)

	asserts.NotNil(db)
}
