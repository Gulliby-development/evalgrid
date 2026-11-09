package main

import "testing"

var rowsRouge = []row{
	{[]string{"score", "--metric", "rouge"}, "ref q1 cat sat\nans q1 m cat sat\n", 0, "m samples 1 cases 1 rouge 1.0000\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 cat sat\nans q1 m sat cat\n", 0, "m samples 1 cases 1 rouge 0.5000\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 a b c d\nans q1 m b d\n", 0, "m samples 1 cases 1 rouge 0.8000\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 x y z\nans q1 m z y x\n", 0, "m samples 1 cases 1 rouge 0.3333\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 x y z w\nans q1 m y x w z\n", 0, "m samples 1 cases 1 rouge 0.5000\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 aa bb aa bb\nans q1 m bb aa bb\n", 0, "m samples 1 cases 1 rouge 0.8571\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 aa bb cc dd\nans q1 m cc aa dd bb\n", 0, "m samples 1 cases 1 rouge 0.5000\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 aa\nans q1 m aa aa aa\n", 0, "m samples 1 cases 1 rouge 0.5000\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 the\nans q1 m\nans q1 m aa\n", 0, "m samples 2 cases 1 rouge 0.5000\n", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 aa bb cc\nans q1 m aa cc\nans q1 m bb\nans q1 m cc bb aa\n", 0, "m samples 3 cases 1 rouge 0.5444\n", ""},
}

func TestRowsRouge(t *testing.T) { checkRows(t, rowsRouge) }
