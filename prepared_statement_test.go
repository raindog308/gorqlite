package gorqlite

import "testing"

func TestPreparedStatement(t *testing.T) {
	cases := []struct {
		input  string
		args   []interface{}
		output string
	}{
		{
			input:  "SELECT * FROM posts WHERE creator=%d",
			args:   []interface{}{42},
			output: "SELECT * FROM posts WHERE creator=42",
		},
		{
			input:  "INSERT INTO posts(body) VALUES(%s)",
			args:   []interface{}{`foo "bar" baz`},
			output: `INSERT INTO posts(body) VALUES('foo \"bar\" baz')`,
		},
	}

	for _, cs := range cases {
		t.Run(cs.input, func(t *testing.T) {
			p := NewPreparedStatement(cs.input)
			outp := p.Bind(cs.args...)

			if outp != cs.output {
				t.Fatalf("expected output to be %s but got: %s", cs.output, outp)
			}
		})
	}
}
