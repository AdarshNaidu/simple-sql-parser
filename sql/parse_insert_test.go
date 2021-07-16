package sql_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/AdarshNaidu/simple-sql-parser/sql"
)

// Ensure the parser can parse strings into Statement ASTs.
func TestParser_ParseInsertStatement(t *testing.T) {
	var tests = []struct {
		s    string
		stmt *sql.InsertStatement
		err  string
	}{
		{
			s: `INSERT INTO employee (name, role) VALUES (kattappa, senapati)`,
			stmt: &sql.InsertStatement{
				Fields: []string{"name", "role"},
				Values: []string{"kattappa", "senapati"},
				TableName: "employee",
			},
		},

		{
			s: `insert into my_table values (hello, world)`,
			stmt: &sql.InsertStatement{
				Fields: nil,
				Values: []string{"hello", "world"},
				TableName: "my_table",
			},
		},

		{
			s: `insert into my_table () values (hello, world)`,
			stmt: &sql.InsertStatement{
				Fields: nil,
				Values: []string{"hello", "world"},
				TableName: "my_table",
			},
		},

		// Errors
		{s: `INSERT INTO my_table (one values (hello, world)`, err: `found "values", expected )`},
		{s: `INSERT INTO my_table (hello, world)`, err: `found "", expected VALUES`},
		{s: `INSERT INTO my_table hello`, err: `found "hello", expected ( or VALUES`},
		{s: `INSERT my_table values (hello)`, err: `found "my_table", expected INTO`},
	}

	for i, tt := range tests {
		stmt, err := sql.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}
