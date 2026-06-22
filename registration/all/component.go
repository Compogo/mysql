package all

import (
	"strings"

	"github.com/Compogo/compogo"
	dbClient "github.com/Compogo/db-client"
	dbMigrator "github.com/Compogo/db-migrator"
	dbSqlGenerator "github.com/Compogo/db-sql-generator"
	"github.com/Compogo/mysql"
	"github.com/Compogo/mysql/registration/manager"
	"github.com/Compogo/mysql/registration/migrator"
	sqlGenerator "github.com/Compogo/mysql/registration/sql_generator"
)

// Component — компонент, объединяющий все регистрации для MySQL.
// Подключает MySQL как драйвер для:
//   - db-client (менеджер БД)
//   - db-migrator (миграции)
//   - db-sql-generator (SQL-генератор)
//
// Пример:
//
//	app.AddComponents(&all.Component)
var Component = &compogo.Component{
	Dependencies: compogo.Components{
		&manager.Component,
		&migrator.Component,
		&sqlGenerator.Component,
		&mysql.Component,
	},
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(func(managerCfg *dbClient.Config, migratorCfg *dbMigrator.Config, generatorCfg *dbSqlGenerator.Config) {
			if strings.ToLower(managerCfg.Driver) != strings.ToLower(mysql.DriverName) {
				return
			}

			migratorCfg.Driver = mysql.DriverName
			generatorCfg.Driver = mysql.DriverName
		})
	}),
}
