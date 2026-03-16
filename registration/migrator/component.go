package migrator

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	migrator "github.com/Compogo/db-migrator"
	"github.com/Compogo/mysql"
	"github.com/golang-migrate/migrate/v4/database"
	migratorMySQL "github.com/golang-migrate/migrate/v4/database/mysql"
)

// Component registers MySQL migration driver with the migrator.
// It depends on both migrator.Component and mysql.Component
// to ensure they are initialized before registration.
//
// Usage:
//
//	compogo.WithComponents(
//	    db_client.Component,
//	    db_migrator.Component,
//	    mysql.Component,
//	    migrator.Component,  // registers MySQL migration driver
//	)
var Component = &component.Component{
	Dependencies: component.Components{
		migrator.Component,
		mysql.Component,
	},
}

// init registers MySQL migration driver constructor with the migrator.
// The constructor extracts the MySQL client from container and creates
// a golang-migrate driver instance.
func init() {
	migrator.Registration(mysql.MySQL, func(container container.Container) (database.Driver, error) {
		var c mysql.Client

		err := container.Invoke(func(mysqlClient mysql.Client) {
			c = mysqlClient
		})
		if err != nil {
			return nil, err
		}

		return migratorMySQL.WithInstance(c.SQL(), &migratorMySQL.Config{})
	})
}
