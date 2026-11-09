package main

import "testing"

var rowsText = []row{
	{[]string{"score", "--metric", "em"}, "ref q1 cat\r\nans q1 m cat\r\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 cat\nans q1 m cat", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "\n\n\nref q1 cat\n\n\nans q1 m cat\n\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "  ref   q1\tcat  \n\tans q1\tm   cat\t\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "# first\n   # second\nref q1 cat\nans q1 m cat # tail\n", 0, "m samples 1 cases 1 em 0.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 # cat\nans q1 m # cat\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 cat\nans q1 m\n", 0, "m samples 1 cases 1 em 0.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 cat\nans q1 m cat\r\r\n", 1, "", "evalgrid: line 2: bad character\n"},
	{[]string{"score"}, "# caf\u00e9\nref q1 x\n", 1, "", "evalgrid: line 1: bad character\n"},
	{[]string{"score"}, "ref q1 x\nans q1 m b\x01d\n", 1, "", "evalgrid: line 2: bad character\n"},
	{[]string{"score"}, "ref q1 x\nans q1 m b\tc\n", 0, "m samples 1 cases 1 em 0.0000 f1 0.0000 rouge 0.0000\n", ""},
}

func TestRowsText(t *testing.T) { checkRows(t, rowsText) }
