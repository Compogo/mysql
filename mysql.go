package mysql

import (
	"database/sql"

	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/db-client/client"
	"github.com/Compogo/db-client/connection"
	logger2 "github.com/Compogo/db-client/logger"
	"github.com/go-sql-driver/mysql"
)

// MySQL is the driver identifier for MySQL database.
// It is used across all Compogo database components (client, migrator, generator).
const MySQL = "mysql"

// Client is an alias for client.Client to provide a cleaner API.
// Users can work with mysql.Client without importing the internal client package.
type Client client.Client

// mysqlClient is the internal implementation of the MySQL client.
// It embeds *sql.DB and provides the required methods.
type mysqlClient struct {
	*sql.DB
}

// NewMySQL creates a new MySQL client with automatic decorators.
// The returned client is wrapped with:
//   - connection.Limiter (circuit breaker for error protection)
//   - logger.Logger (query logging at DEBUG level)
//
// This provides production-ready behavior out of the box.
func NewMySQL(logger logger.Logger, config *Config) (Client, error) {
	mysqlConfig, err := mysql.ParseDSN(config.DSN)
	if err != nil {
		return nil, err
	}

	connector, err := mysql.NewConnector(mysqlConfig)
	if err != nil {
		return nil, err
	}

	return logger2.NewLogger(
		connection.NewLimiter(
			&mysqlClient{sql.OpenDB(connector)},
			int64(config.ErrorLimit),
			config.ErrorTimeout,
		),
		logger,
	), nil
}

func (m *mysqlClient) SQL() *sql.DB {
	return m.DB
}

func (m *mysqlClient) DriverName() string {
	return MySQL
}
