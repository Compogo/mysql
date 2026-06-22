package migrator

import (
	"fmt"

	"github.com/Compogo/compogo"
	migrator "github.com/Compogo/db-migrator"
	"github.com/Compogo/mysql"
	"github.com/golang-migrate/migrate/v4/database"
	migratorMySQL "github.com/golang-migrate/migrate/v4/database/mysql"
)

const migrationsTableFormat = "migrations_%s"

// Component — компонент регистрации MySQL драйвера для мигратора.
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&migrator.Component,
		&mysql.Component,
	},
}

// Регистрация драйвера MySQL в системе мигратора.
func init() {
	migrator.Registration(mysql.DriverName, func(container compogo.Container) (database.Driver, error) {
		var client mysql.Client
		var config *compogo.Config

		err := container.Invoke(func(c mysql.Client) {
			client = c
		})
		if err != nil {
			return nil, err
		}

		err = container.Invoke(func(c *compogo.Config) {
			config = c
		})
		if err != nil {
			return nil, err
		}

		return migratorMySQL.WithInstance(client.SQL(), &migratorMySQL.Config{MigrationsTable: fmt.Sprintf(migrationsTableFormat, config.Name)})
	})
}
