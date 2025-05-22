package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

var dsn string

func Init() {
	dsn = viper.GetString("DB_DSN")
}

func InsertTemperatureRow(ctx context.Context, tempC float64, timestamp time.Time) error {
	// Check if time is zero
	if timestamp.IsZero() {
		timestamp = time.Now()
	}

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	stmtName := "insert_temperature_row"
	sql := `
        INSERT INTO temperatures (temperature_c, timestamp)
				VALUES ($1, $2)
				RETURNING id, temperature_c, timestamp;
    `
	if _, err := conn.Prepare(ctx, stmtName, sql); err != nil {
		return err
	}

	if _, err := conn.Exec(ctx, stmtName, tempC, timestamp); err != nil {
		return err
	}

	return nil
}
