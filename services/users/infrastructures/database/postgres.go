package infrastructures

import (
	"time"

	"github.com/adhitamafikri/go-simple-pms/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConn(dsn string) *sqlx.DB {
	log := logger.NewLogger("NewPostgresConn")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal().Msg("Failed to connect to the postgres")
	}

	// connection pool tuning
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
