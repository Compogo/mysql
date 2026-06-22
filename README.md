# Compogo MySQL

Адаптер [MySQL](https://www.mysql.com/) для фреймворка [Compogo](https://github.com/Compogo/compogo).

Предоставляет:

* Клиент для работы с MySQL
* Встроенный Limiter для защиты от ошибок соединения
* Логирование всех SQL-запросов
* Интеграцию с `db-client` как драйвер
* Интеграцию с `db-migrator` для миграций
* Интеграцию с `db-sql-generator` для генерации SQL

## Установка

```shell
go get github.com/Compogo/mysql
```

## Быстрый старт

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/mysql/registration/all"
)

func main() {
    app := compogo.NewApp("myapp",
        // Подключаем всё сразу (клиент + мигратор + генератор)
        compogo.WithComponents(&all.Component),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

## Конфигурация

### Флаги командной строки

```shell
# DSN строка подключения
--db.mysql.dsn="user:pass@tcp(localhost:3306)/dbname?parseTime=true"

# Лимит ошибок соединения (защита от "штурма")
--db.mysql.connection.error.limit=10

# Время блокировки при превышении лимита
--db.mysql.connection.error.timeout=10s
```

### Пример DSN

```shell
# Базовая строка
user:password@tcp(localhost:3306)/dbname

# С параметрами
user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=true&loc=UTC

# Без пароля
user@tcp(localhost:3306)/dbname
```

## Защита от ошибок соединения

При возникновении ошибок соединения (sql.ErrConnDone, driver.ErrBadConn) клиент:

1. Увеличивает счётчик ошибок.
2. При превышении лимита (по умолчанию 10) блокирует запросы на указанное время.
3. После таймаута восстанавливает работу.

Это защищает БД от "штурма" при сетевых проблемах.

## Регистрация компонентов

### Всё сразу

```go
import "github.com/Compogo/mysql/registration/all"

app.AddComponents(&all.Component)
```

### По отдельности

```go
import (
    "github.com/Compogo/mysql"
    "github.com/Compogo/mysql/registration/manager"
    "github.com/Compogo/mysql/registration/migrator"
    "github.com/Compogo/mysql/registration/sql_generator"
)

// Только клиент БД
app.AddComponents(&mysql.Component)

// Клиент + регистрация в db-client
app.AddComponents(&manager.Component)

// Клиент + регистрация в db-migrator
app.AddComponents(&migrator.Component)

// Клиент + регистрация в db-sql-generator
app.AddComponents(&sqlGenerator.Component)
```

## Зависимости

* [Compogo](https://github.com/Compogo/compogo) — основной фреймворк
* [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) — драйвер MySQL
* [Compogo DB Client](https://github.com/Compogo/db-client) — клиент БД
* [Compogo DB Migrator](https://github.com/Compogo/db-migrator) — мигратор
* [Compogo DB SQL Generator](https://github.com/Compogo/db-sql-generator) — SQL-генератор

## Лицензия

```plantuml
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
