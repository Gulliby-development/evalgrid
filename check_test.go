package main

import "testing"

var rowsCheck = []row{
	{[]string{"check"}, "ref q1 yes\nans q1 m yes\n", 0, "m em 100.00% ok\n", ""},
	{[]string{"check"}, "ref q1 yes\nans q1 m no\n", 3, "m em 0.00% below\n", ""},
	{[]string{"check"}, "ref q1 yes\nans q1 m yes\nans q1 m no\n", 0, "m em 50.00% ok\n", ""},
	{[]string{"check"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m no\n", 3, "m em 33.33% below\n", ""},
	{[]string{"check", "--min", "75"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m yes\n", 3, "m em 66.67% below\n", ""},
	{[]string{"check", "--min", "100"}, "ref q1 yes\nans q1 m yes\nans q1 m no\n", 3, "m em 50.00% below\n", ""},
	{[]string{"check", "--min", "007"}, "ref q1 yes\nans q1 m yes\nans q1 m no\n", 0, "m em 50.00% ok\n", ""},
	{[]string{"check"}, "", 0, "no answers\n", ""},
	{[]string{"check"}, "ref q1 yes\nans q1 good yes\nans q1 bad no\nans q1 mid yes\nans q1 mid no\n", 3, "bad em 0.00% below\ngood em 100.00% ok\nmid em 50.00% ok\n", ""},
	{[]string{"check", "--min", "33"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m no\n", 0, "m em 33.33% ok\n", ""},
	{[]string{"check"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m no\nans q1 m no\nans q1 m no\nans q1 m no\nans q1 m no\n", 3, "m em 14.29% below\n", ""},
}

func TestRowsCheck(t *testing.T) { checkRows(t, rowsCheck) }
