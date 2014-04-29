# Gorel
![Release](http://img.shields.io/github/release/rastasheep/gorel.svg?style=flat)

Minimalistic framework which allows you to construct and manage coplex SQL queries.

__NOTE:__ This is pre-alpha software.

It's not likely that it could help you with developing your application. However if you are interesting in learning Golang or just want to contribute you came to the right place.

## Features
* **Fully compatible with the [database/sql](http://godoc.org/database/sql) package.**
* Thin layer around SQL queries.
* Extremely simple to use.
* Non-intrusive design.
* Plays nice with all Golang packages / frameworks.
* Intended to be a framework framework, works standalone too.

## A Gentle Introduction

Generating a basic query is very simple. For example, in order to produce

```sql
SELECT * FROM users
```

you construct a table relation and convert it to sql:

```go
users = gorel.Table("users")
query = users.Select("*")
query.ToSql
```

```go
query = users.where(users.AttrEqual("name", "amy") )

# => SELECT * FROM users WHERE users.name = "amy"
```


## About

Inspired by [Arel](https://github.com/rails/arel) and [Ecto](https://github.com/elixir-lang/ecto)

Gorel is created and designed by [rastasheep](https://github.com/rastasheep/) and it's released under the [MIT License](http://opensource.org/licenses/MIT).
