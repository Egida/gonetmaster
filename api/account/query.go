package account

import (
	"context"
	"time"

	"github.com/edwardelton/gonetmaster/api/config"
	"github.com/edwardelton/gonetmaster/api/database"
)

func RetrieveApiKey(hashedKey string) (bool, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.ConnectionTimeout*time.Second)
	defer cancelFunc()

	rows, queryErr := database.Db.QueryContext(ctx, "SELECT key FROM account WHERE key = $1", hashedKey)
	defer rows.Close()

	if queryErr != nil {
		return false, queryErr
	}

	if !rows.Next() {
		return false, nil
	}
	return true, nil
}
