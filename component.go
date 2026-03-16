package mysql

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/flag"
)

// Component is a ready-to-use Compogo component that provides a MySQL client.
// It automatically:
//   - Registers Config and NewMySQL in the DI container
//   - Adds command-line flags for MySQL configuration
//   - Applies configuration during Configuration phase
//
// Usage:
//
//	compogo.WithComponents(
//	    db_client.Component,
//	    mysql.Component,
//	    // ... other components
//	)
//
// The MySQL client can be injected into any component:
//
//	type UserService struct {
//	    db mysql.Client
//	}
var Component = &component.Component{
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provides(
			NewConfig,
			NewMySQL,
		)
	}),
	BindFlags: component.BindFlags(func(flagSet flag.FlagSet, container container.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.DSN, DsnFieldName, "", "mysql dsn string connection")
			flagSet.Uint32Var(&config.ErrorLimit, ConnectionErrorLimitFieldName, ConnectionErrorLimitDefault, "")
			flagSet.DurationVar(&config.ErrorTimeout, ConnectionErrorTimeoutFieldName, ConnectionErrorTimeoutDefault, "")
		})
	}),
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(Configuration)
	}),
}
