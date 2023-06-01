package account

import (
	"context"
	"time"

	"github.com/edwardelton/gonetmaster/api/config"
	"github.com/edwardelton/gonetmaster/api/database"
	"github.com/edwardelton/gonetmaster/logger"
)

func RetrieveApiKey(hashedKey string) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.ConnectionTimeout*time.Second)
	defer cancelFunc()

	rows, queryErr := database.Db.QueryContext(ctx, "SELECT key FROM account WHERE key = $1", hashedKey)
	defer rows.Close()

	if queryErr != nil {
		logger.Log.Warn("Error retrieving apiKey: ", queryErr)
		return false, queryErr
	}

	if !rows.Next() {
		logger.Log.Warn("No ApiKey found")
		return false, nil
	}
	return true, nil
}
