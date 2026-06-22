package mysql

import (
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
	"github.com/Compogo/runner"
)

// Component — компонент MySQL для Compogo.
// Регистрирует конфигурацию и клиент в DI-контейнере.
//
// Пример подключения:
//
//	app.AddComponents(&mysql.Component)
//
//	var client mysql.Client
//	container.Invoke(func(c mysql.Client) { client = c })
//	rows, err := client.Query("SELECT * FROM users")
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&runner.Component,
	},
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provides(
			NewConfig,
			NewMySQL,
		)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.DSN, DsnFieldName, "", "mysql dsn string connection")
			flagSet.Uint64Var(&config.ErrorLimit, ConnectionErrorLimitFieldName, ConnectionErrorLimitDefault, "")
			flagSet.DurationVar(&config.ErrorTimeout, ConnectionErrorTimeoutFieldName, ConnectionErrorTimeoutDefault, "")
		})
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
}
