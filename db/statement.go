package db

import (
	"fmt"
	"strings"
)

type statement struct {
	sql string
}

func newstatement() *statement {
	stmt := new(statement)

	return stmt
}

func (stmt *statement) selects(fields ...string) *statement {
	stmt.sql = fmt.Sprintf("SELECT %s", strings.Join(fields, ","))
	return stmt
}

// Count sql count
func (stmt *statement) count(field string) *statement {
	stmt.sql = fmt.Sprintf("SELECT COUNT(%s)", field)
	return stmt
}

// From sql from
func (stmt *statement) from(table string) *statement {
	sql := fmt.Sprintf(" FROM %s", table)
	stmt.sql += sql
	return stmt
}

// Where sql where
func (stmt *statement) where(field string, value string, or bool) *statement {
	sqlCond := fmt.Sprintf("%s = '%s'", field, value)
	whereStr := "WHERE"
	if strings.Contains(stmt.sql, whereStr) {
		if or {
			whereStr = "OR"
		} else {
			whereStr = "AND"
		}
	}

	stmt.sql = fmt.Sprintf("%s %s %s", stmt.sql, whereStr, sqlCond)
	return stmt
}

// Insert sql insert
func (stmt *statement) insert(table string) *statement {
	sql := fmt.Sprintf("INSERT INTO %s", table)
	stmt.sql = sql

	return stmt
}

// Columns sql insert columns
func (stmt *statement) columns(fields ...string) *statement {
	if len(fields) > 0 {
		stmt.sql = fmt.Sprintf("%s (%s)", stmt.sql, strings.Join(fields, ","))
	}
	return stmt
}

// Values sql insert values
func (stmt *statement) values(vals ...string) *statement {
	for i, s := range vals {
		vals[i] = strings.Replace(s, "'", "\\'", -1)
	}
	s := strings.Join(vals, "','")
	s = "'" + s + "'"
	stmt.sql = fmt.Sprintf("%s VALUES (%s)", stmt.sql, s)
	return stmt
}

func (stmt *statement) orderBy(col string, asc bool) *statement {
	order := "DESC"
	if asc {
		order = "ASC"
	}
	stmt.sql = fmt.Sprintf("%s ORDER BY %s %s", stmt.sql, col, order)
	return stmt
}

func (stmt *statement) limit(n ...int) *statement {
	stmt.sql = fmt.Sprintf("%s LIMIT %s",
		stmt.sql,
		strings.Trim(strings.Replace(fmt.Sprint(n), " ", ",", -1), "[]"))
	return stmt
}

// SQL return sql statement string
func (stmt *statement) toString() string {
	return stmt.sql + ";"
}
