package manager

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	dbClient "github.com/Compogo/db-client"
	"github.com/Compogo/db-client/client"
	"github.com/Compogo/mysql"
)

// Component registers MySQL client factory with db-client.
// It depends on both dbClient.Component and mysql.Component
// to ensure they are initialized before registration.
//
// Usage:
//
//	compogo.WithComponents(
//	    db_client.Component,
//	    mysql.Component,
//	    manager.Component,  // registers MySQL client factory
//	)
var Component = &component.Component{
	Dependencies: component.Components{
		dbClient.Component,
		mysql.Component,
	},
}

// init registers MySQL client constructor with db-client.
// The constructor retrieves the MySQL client from container and
// returns it as a client.Client interface.
func init() {
	dbClient.Registration(mysql.MySQL, func(container container.Container) (client.Client, error) {
		var c mysql.Client

		err := container.Invoke(func(mysqlClient mysql.Client) {
			c = mysqlClient
		})
		if err != nil {
			return nil, err
		}

		return c, nil
	})
}
