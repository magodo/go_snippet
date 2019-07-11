package dump

import "testing"

func TestFoo(t *testing.T) {
	cases := []struct {
		query string
		args  []interface{}
		dump  string
	}{
		{
			`$$,$1,$2,$1,$3`,
			[]interface{}{"a", 123, "$"},
			`$$,'a',123,'a','$'`,
		},
		{
			`$1,$10`,
			[]interface{}{"a", 2, 3, 4, 5, 6, 7, 8, 9, "b"},
			`'a','b'`,
		},
		{
			`select * from (select * from a where id = $1)`,
			[]interface{}{"a"},
			`select * from (select * from a where id = 'a')`,
		},
	}

	for _, c := range cases {
		out, err := showPostgreSQLStmt(c.query, c.args...)
		if err != nil {
			t.Fatal(err)
		}
		if out != c.dump {
			t.Fatalf("expect: %s\nbut got: %s", c.dump, out)
		}
	}
}
