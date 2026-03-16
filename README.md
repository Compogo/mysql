# Compogo MySQL 🐬

**Compogo MySQL** — это полноценный MySQL-драйвер для экосистемы Compogo. Включает клиент с автоматической защитой от сбоев, поддержку миграций, генератор SQL-запросов и готовые регистрации для всех компонентов.

## 🚀 Установка

```bash
go get github.com/Compogo/mysql
```
### 📦 Быстрый старт (одна БД, всё автоматически)

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/db-client"
    "github.com/Compogo/db-migrator"
    "github.com/Compogo/db-sql-generator"
    "github.com/Compogo/mysql"
    "github.com/Compogo/mysql/registration/all"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithOsSignalCloser(),
        db_client.Component,
        db_migrator.Component,
        db_sql_generator.Component,
        mysql.Component,
        all.Component,  // ← магия: всё настраивается само
        compogo.WithComponents(
            userRepositoryComponent,
        ),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}

// Репозиторий работает с общим интерфейсом
var userRepositoryComponent = &component.Component{
    Dependencies: component.Components{db_client.Component},
    Execute: component.StepFunc(func(c container.Container) error {
        return c.Invoke(func(db db_client.Client) {
            repo := &UserRepository{db: db}
            // ...
        })
    }),
}
```

Запуск:

```shell
./myapp --db.driver=mysql --db.mysql.dsn="user:pass@tcp(localhost:3306)/db"
```

### ✨ Возможности

#### 🛡️ Клиент с защитой из коробки

Каждый MySQL-клиент автоматически оборачивается:

* Circuit breaker — защита от временных сбоев (настраиваемые лимит и таймаут)
* Логирование — все запросы логируются на уровне DEBUG
