package mysql

import (
	"time"

	"github.com/Compogo/compogo"
)

const (
	// DsnFieldName — имя поля для DSN строки подключения.
	DsnFieldName = "db.mysql.dsn"

	// ConnectionErrorLimitFieldName — имя поля для лимита ошибок соединения.
	ConnectionErrorLimitFieldName = "db.mysql.connection.error.limit"

	// ConnectionErrorTimeoutFieldName — имя поля для таймаута при ошибках соединения.
	ConnectionErrorTimeoutFieldName = "db.mysql.connection.error.timeout"
)

var (
	// ConnectionErrorLimitDefault — лимит ошибок соединения по умолчанию.
	ConnectionErrorLimitDefault = uint64(10)

	// ConnectionErrorTimeoutDefault — таймаут при ошибках соединения по умолчанию.
	ConnectionErrorTimeoutDefault = 10 * time.Second
)

// Config содержит конфигурацию MySQL.
type Config struct {
	DSN          string
	ErrorLimit   uint64
	ErrorTimeout time.Duration
}

// NewConfig создаёт новую конфигурацию.
func NewConfig() *Config {
	return &Config{}
}

// Configuration загружает конфигурацию из Configurator.
// Если DSN не задан, он должен быть получен из Configurator.
func Configuration(config *Config, configurator compogo.Configurator) *Config {
	if config.DSN == "" {
		config.DSN = configurator.GetString(DsnFieldName)
	}

	if config.ErrorLimit == 0 || config.ErrorLimit == ConnectionErrorLimitDefault {
		configurator.SetDefault(ConnectionErrorLimitFieldName, ConnectionErrorLimitDefault)
		config.ErrorLimit = configurator.GetUint64(ConnectionErrorLimitFieldName)
	}

	if config.ErrorTimeout == 0 || config.ErrorTimeout == ConnectionErrorTimeoutDefault {
		configurator.SetDefault(ConnectionErrorTimeoutFieldName, ConnectionErrorTimeoutDefault)
		config.ErrorTimeout = configurator.GetDuration(ConnectionErrorTimeoutFieldName)
	}

	return config
}
