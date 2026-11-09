package main

import "testing"

type row struct {
	args   []string
	in     string
	status int
	out    string
	err    string
}

func checkRows(t *testing.T, rows []row) {
	t.Helper()
	for i, r := range rows {
		status, out, errText := run(r.args, r.in)
		if status != r.status || out != r.out || errText != r.err {
			t.Errorf("row %d %q: got (%d, %q, %q), want (%d, %q, %q)", i+1, r.args, status, out, errText, r.status, r.out, r.err)
		}
	}
}
