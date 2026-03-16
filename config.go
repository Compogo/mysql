package mysql

import (
	"time"

	"github.com/Compogo/compogo/configurator"
)

const (
	DsnFieldName = "db.mysql.dsn"

	ConnectionErrorLimitFieldName   = "db.mysql.connection.error.limit"
	ConnectionErrorTimeoutFieldName = "db.mysql.connection.error.timeout"

	ConnectionErrorLimitDefault   = uint32(10)
	ConnectionErrorTimeoutDefault = 10 * time.Second
)

type Config struct {
	DSN          string
	ErrorLimit   uint32
	ErrorTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{}
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	if config.DSN == "" {
		config.DSN = configurator.GetString(DsnFieldName)
	}

	if config.ErrorLimit == 0 || config.ErrorLimit == ConnectionErrorLimitDefault {
		configurator.SetDefault(ConnectionErrorLimitFieldName, ConnectionErrorLimitDefault)
		config.ErrorLimit = configurator.GetUint32(ConnectionErrorLimitFieldName)
	}

	if config.ErrorTimeout == 0 || config.ErrorTimeout == ConnectionErrorTimeoutDefault {
		configurator.SetDefault(ConnectionErrorTimeoutFieldName, ConnectionErrorTimeoutDefault)
		config.ErrorTimeout = configurator.GetDuration(ConnectionErrorTimeoutFieldName)
	}

	return config
}
