package gorqlite

import (
	"fmt"
	"strings"
)

// EscapeString sql-escapes a string.
func EscapeString(value string) string {
	replace := [][2]string{
		{`\`, `\\`},
		{`\0`, `\\0`},
		{`\n`, `\\n`},
		{`\r`, `\\r`},
		{`"`, `\"`},
		{`'`, `\'`},
	}

	for _, val := range replace {
		value = strings.Replace(value, val[0], val[1], -1)
	}

	return value
}

// PreparedStatement is a simple wrapper around fmt.Sprintf for prepared SQL
// statements.
type PreparedStatement struct {
	body string
}

// NewPreparedStatement takes a sprintf syntax SQL query for later binding of
// parameters.
func NewPreparedStatement(body string) PreparedStatement {
	return PreparedStatement{body: body}
}

// Bind takes arguments and SQL-escapes them, then calling fmt.Sprintf.
func (p PreparedStatement) Bind(args ...interface{}) string {
	var spargs []interface{}

	for _, arg := range args {
		switch arg.(type) {
		case string:
			spargs = append(spargs, `'`+EscapeString(arg.(string))+`'`)
		case fmt.Stringer:
			spargs = append(spargs, `'`+EscapeString(arg.(fmt.Stringer).String())+`'`)
		default:
			spargs = append(spargs, arg)
		}
	}

	return fmt.Sprintf(p.body, spargs...)
}
