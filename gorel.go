package gorel

import "fmt"

type GorelTable struct {
	name string
	pick string
}

func Table(name string) *GorelTable {
	t := &GorelTable{name: name, pick: "*"}

	return t
}

func (table *GorelTable) Select(qury string) *GorelTable {
	table.pick = qury

	return table
}

func (table *GorelTable) ToSql() string {
	query := fmt.Sprint("SELECT ", table.pick, " FROM ", table.name, ";")

	return query
}
