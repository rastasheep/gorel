package gorel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTable(t *testing.T) {
	table := Table("users")

	assert.Equal(t, table.name, "users", "table name should be set")
	assert.Equal(t, table.pick, "*", "default select params should be set")
}

func TestSelect(t *testing.T) {
	table := Table("users")
	table.Select("name")

	assert.Equal(t, table.pick, "name", "select  should change tables param")
}

func TestToSql(t *testing.T) {
	query := Table("users").ToSql()
	assert.Equal(t, query, "SELECT * FROM users;", "it should create default select query")

	query = Table("users").Select("name").ToSql()
	assert.Equal(t, query, "SELECT name FROM users;", "it should create select query with params")

	query = Table("users").Select("name, age").ToSql()
	assert.Equal(t, query, "SELECT name, age FROM users;", "it should create select query with params")
}
