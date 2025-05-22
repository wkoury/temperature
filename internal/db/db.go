package db

import 	(
	"github.com/spf13/viper"
	    "github.com/jackc/pgx/v5"
)

var DSN string

func Init() {
DSN = viper.GetString("DB_DSN")
}

func InsertTemperatureRow(ctx context.Context, tempC float64, time time.Time) error {
	if time == time.Time{} {
		time = time.Now()
	}
}