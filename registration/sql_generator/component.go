package sql_generator

import (
	"github.com/Compogo/compogo"
	sqlGenerator "github.com/Compogo/db-sql-generator"
	"github.com/Compogo/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

const DialectName = "mysql"

// Component — компонент регистрации MySQL диалекта для SQL-генератора.
// Регистрирует соответствие между драйвером "mysql" и диалектом "mysql".
var Component = compogo.Component{
	Dependencies: compogo.Components{
		&sqlGenerator.Component,
		&mysql.Component,
	},
}

// Регистрация диалекта MySQL в системе SQL-генератора.
func init() {
	sqlGenerator.Registration(mysql.DriverName, DialectName)
}
