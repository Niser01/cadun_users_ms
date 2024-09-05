package DB

import (
	"context"
	"fmt"

	"github.com/Niser01/CADUN_Users/tree/main/cadun_users_ms/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// New returns a new sqlx.DB object with the connection to the PostgreSQL database
func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	// Formato de cadena de conexión para PostgreSQL, con SSL habilitado
	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=require", // Cambiado sslmode a 'require'
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)

	// Cambia "mysql" a "postgres"
	return sqlx.ConnectContext(ctx, "postgres", connectionString)
}
