package sql_generator

import (
	"github.com/Compogo/compogo/component"
	sqlGenerator "github.com/Compogo/db-sql-generator"
	"github.com/Compogo/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

// Component registers MySQL dialect with the SQL generator.
// It depends on both sqlGenerator.Component and mysql.Component
// to ensure they are initialized before registration.
//
// Usage:
//
//	compogo.WithComponents(
//	    db_client.Component,
//	    db_sql_generator.Component,
//	    mysql.Component,
//	    sql_generator.Component,  // registers MySQL dialect
//	)
var Component = &component.Component{
	Dependencies: component.Components{
		sqlGenerator.Component,
		mysql.Component,
	},
}

// init registers MySQL dialect alias with the SQL generator.
// This allows goqu to generate MySQL-specific SQL syntax.
func init() {
	sqlGenerator.Registration(mysql.MySQL, "mysql")
}
