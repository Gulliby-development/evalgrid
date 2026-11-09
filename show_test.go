package main

import "testing"

var rowsShow = []row{
	{[]string{"show", "q1"}, "ref q1 cat sat\nans q1 m cat sat\n", 0, "ref cat sat\nm em 1 f1 1.0000 rouge 1.0000\n", ""},
	{[]string{"show", "q1"}, "ref q1 The Cat sat!\nans q1 m cat\nans q1 n sat cat\n", 0, "ref The Cat sat!\nm em 0 f1 0.6667 rouge 0.6667\nn em 0 f1 1.0000 rouge 0.5000\n", ""},
	{[]string{"show", "q2"}, "ref q1 x\nref q2 y\nans q1 m x\nans q2 m y\n", 0, "ref y\nm em 1 f1 1.0000 rouge 1.0000\n", ""},
	{[]string{"show", "q1"}, "ref q1 x\n", 0, "ref x\n", ""},
	{[]string{"show", "q9"}, "ref q1 x\n", 1, "", "evalgrid: no case q9\n"},
	{[]string{"show", "q1"}, "", 1, "", "evalgrid: no case q1\n"},
	{[]string{"show", "q1"}, "ref q1 x\nans q1 m y\nans q2 m z\n", 1, "", "evalgrid: line 3: no reference for q2\n"},
	{[]string{"show", "q1"}, "ref q1   Spaced   out  \nans q1 m spaced\n", 0, "ref Spaced   out\nm em 0 f1 0.6667 rouge 0.6667\n", ""},
	{[]string{"show", "q1"}, "ref q1 aa bb cc\nans q1 m cc aa\nans q1 m\nans q1 m aa bb cc dd\n", 0, "ref aa bb cc\nm em 0 f1 0.8000 rouge 0.4000\nm em 0 f1 0.0000 rouge 0.0000\nm em 0 f1 0.8571 rouge 0.8571\n", ""},
	{[]string{"show", "q1"}, "ref q1 the\nans q1 m\nans q1 m an\n", 0, "ref the\nm em 1 f1 1.0000 rouge 1.0000\nm em 1 f1 1.0000 rouge 1.0000\n", ""},
	{[]string{"show", "q1"}, "ref q1 x\nans q1 b x\nans q1 a y\nans q1 b y\n", 0, "ref x\nb em 1 f1 1.0000 rouge 1.0000\na em 0 f1 0.0000 rouge 0.0000\nb em 0 f1 0.0000 rouge 0.0000\n", ""},
}

func TestRowsShow(t *testing.T) { checkRows(t, rowsShow) }
