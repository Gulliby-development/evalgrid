package main

import "testing"

var rowsPassk = []row{
	{[]string{"passk", "1"}, "ref q1 yes\nans q1 m yes\n", 0, "m cases 1 skipped 0 pass@1 1.0000\n", ""},
	{[]string{"passk", "1"}, "ref q1 yes\nans q1 m yes\nans q1 m no\n", 0, "m cases 1 skipped 0 pass@1 0.5000\n", ""},
	{[]string{"passk", "2"}, "ref q1 yes\nans q1 m yes\nans q1 m no\n", 0, "m cases 1 skipped 0 pass@2 1.0000\n", ""},
	{[]string{"passk", "2"}, "ref q1 yes\nans q1 m no\nans q1 m no\n", 0, "m cases 1 skipped 0 pass@2 0.0000\n", ""},
	{[]string{"passk", "3"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m no\nans q1 m no\n", 0, "m cases 1 skipped 0 pass@3 0.7500\n", ""},
	{[]string{"passk", "2"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m no\nans q1 m no\n", 0, "m cases 1 skipped 0 pass@2 0.5000\n", ""},
	{[]string{"passk", "2"}, "ref q1 yes\nref q2 no\nans q1 m yes\nans q1 m yes\nans q2 m yes\nans q2 m no\nans q2 m yes\n", 0, "m cases 2 skipped 0 pass@2 0.8333\n", ""},
	{[]string{"passk", "3"}, "ref q1 yes\nref q2 yes\nans q1 m yes\nans q1 m yes\nans q2 m no\n", 0, "m cases 0 skipped 2 pass@3 -\n", ""},
	{[]string{"passk", "07"}, "ref q1 yes\nans q1 m yes\n", 0, "m cases 0 skipped 1 pass@7 -\n", ""},
	{[]string{"passk", "1"}, "ref q1 yes\nans q1 b Yes!\nans q1 a the yes\nans q1 a no\n", 0, "a cases 1 skipped 0 pass@1 0.5000\nb cases 1 skipped 0 pass@1 1.0000\n", ""},
	{[]string{"passk", "1"}, "", 0, "", ""},
	{[]string{"passk", "5"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m no\nans q1 m yes\nans q1 m no\nans q1 m no\nans q1 m yes\nans q1 m no\nans q1 m no\n", 0, "m cases 1 skipped 0 pass@5 0.9524\n", ""},
	{[]string{"passk", "4"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m no\nans q1 m yes\nans q1 m no\nans q1 m no\nans q1 m yes\nans q1 m no\nans q1 m no\n", 0, "m cases 1 skipped 0 pass@4 0.8810\n", ""},
}

func TestRowsPassk(t *testing.T) { checkRows(t, rowsPassk) }
