package mylib

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		query string
		args  []interface{}
		dump  string
	}{
		{
			`?,?,?`,
			[]interface{}{"a", 123, "$"},
			`'a',123,'$'`,
		},
	}

	for _, c := range cases {
		out, err := mysqlDumpQuery(c.query, c.args...)
		if err != nil {
			t.Fatal(err)
		}
		if out != c.dump {
			t.Fatalf("expect: %s\nbut got: %s", c.dump, out)
		}
	}
}
