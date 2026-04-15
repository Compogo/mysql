package all

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	dbClient "github.com/Compogo/db-client"
	dbMigrator "github.com/Compogo/db-migrator"
	dbSqlGenerator "github.com/Compogo/db-sql-generator"
	"github.com/Compogo/mysql"
	"github.com/Compogo/mysql/registration/manager"
	"github.com/Compogo/mysql/registration/migrator"
	sqlGenerator "github.com/Compogo/mysql/registration/sql_generator"
)

// Component is a complete MySQL integration that automatically:
//   - Provides the MySQL client
//   - Registers MySQL with db-client
//   - Registers MySQL migration driver with db-migrator
//   - Registers MySQL dialect with db-sql-generator
//   - Automatically propagates driver selection to all components
//
// Usage (recommended for single-database applications):
//
//	compogo.WithComponents(
//	    db_client.Component,
//	    db_migrator.Component,
//	    db_sql_generator.Component,
//	    mysql.Component,
//	    all.Component,  // ← everything automatically wired
//	)
//
// Then just select the driver via flag:
//
//	./myapp --db.driver=mysql --db.mysql.dsn="..."
//
// The migrator and SQL generator will automatically use the same driver.
var Component = &component.Component{
	Dependencies: component.Components{
		manager.Component,
		migrator.Component,
		sqlGenerator.Component,
		mysql.Component,
	},
	Configuration: component.StepFunc(func(container container.Container) error {
		return container.Invoke(func(managerCfg *dbClient.Config, migratorCfg *dbMigrator.Config, generatorCfg *dbSqlGenerator.Config) {
			if managerCfg.Driver != mysql.MySQL {
				return
			}

			migratorCfg.Driver = mysql.MySQL
			generatorCfg.Driver = mysql.MySQL
		})
	}),
}
