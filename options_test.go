package main

import "testing"

var rowsOptions = []row{
	{[]string{"score", "--metric"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: --metric needs a value\n"},
	{[]string{"score", "--metric", ""}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: --metric needs a value\n"},
	{[]string{"score", "--metric", "em", "--metric", "f1"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: --metric given twice\n"},
	{[]string{"score", "--metric", "-x"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: bad metric -x\n"},
	{[]string{"score", "--metric", "EM"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: bad metric EM\n"},
	{[]string{"score", "extra"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: unexpected argument extra\n"},
	{[]string{"passk"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: passk needs K\n"},
	{[]string{"passk", "1", "2"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: unexpected argument 2\n"},
	{[]string{"passk", "0"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: bad k 0\n"},
	{[]string{"passk", "1x"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: bad k 1x\n"},
	{[]string{"show"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: show needs ID\n"},
	{[]string{"show", "Q1"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: bad id Q1\n"},
	{[]string{"show", "q1", "q2"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: unexpected argument q2\n"},
	{[]string{"check", "--min"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: --min needs a value\n"},
	{[]string{"check", "--min", "101"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: --min needs an integer from 0 to 100\n"},
	{[]string{"check", "--min", "x"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: --min needs an integer from 0 to 100\n"},
	{[]string{"check", "--min", "5", "--min", "6"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: --min given twice\n"},
	{[]string{"check", "--min", "101", "extra"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: unexpected argument extra\n"},
}

func TestRowsOptions(t *testing.T) { checkRows(t, rowsOptions) }
