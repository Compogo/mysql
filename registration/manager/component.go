package manager

import (
	"github.com/Compogo/compogo"
	dbClient "github.com/Compogo/db-client"
	"github.com/Compogo/mysql"
)

// Component — компонент регистрации MySQL драйвера для db-client.
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&dbClient.Component,
		&mysql.Component,
	},
}

// Регистрация драйвера MySQL в системе db-client.
func init() {
	dbClient.Registration(mysql.DriverName, func(container compogo.Container) (dbClient.Client, error) {
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
