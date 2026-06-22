package mysql

import (
	"database/sql"

	"github.com/Compogo/compogo"
	dbClient "github.com/Compogo/db-client"
	"github.com/Compogo/runner"
	"github.com/go-sql-driver/mysql"
)

// DriverName — имя драйвера MySQL.
const DriverName = "mysql"

// Client — интерфейс клиента MySQL (наследует db_client.Client).
type Client dbClient.Client

// mysqlClient — внутренняя реализация клиента MySQL.
type mysqlClient struct {
	*sql.DB
}

// NewMySQL создаёт новый клиент MySQL.
// Оборачивает соединение в:
//   - Limiter — защита от ошибок соединения
//   - Logger — логирование всех SQL-запросов
func NewMySQL(logger compogo.Logger, config *Config, runner runner.Runner) (Client, error) {
	mysqlConfig, err := mysql.ParseDSN(config.DSN)
	if err != nil {
		return nil, err
	}

	connector, err := mysql.NewConnector(mysqlConfig)
	if err != nil {
		return nil, err
	}

	return dbClient.NewLogger(
		dbClient.NewLimiter(
			&mysqlClient{sql.OpenDB(connector)},
			config.ErrorLimit,
			config.ErrorTimeout,
			runner,
		),
		logger,
	), nil
}

func (m *mysqlClient) SQL() *sql.DB {
	return m.DB
}

func (m *mysqlClient) DriverName() string {
	return DriverName
}
