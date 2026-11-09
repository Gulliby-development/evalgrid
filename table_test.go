package main

import "testing"

var rowsTable = []row{
	{[]string{"score"}, "ref q1 cat sat\nans q1 m cat sat\n", 0, "m samples 1 cases 1 em 1.0000 f1 1.0000 rouge 1.0000\n", ""},
	{[]string{"score"}, "ref q1 cat sat\nans q1 m sat cat on\n", 0, "m samples 1 cases 1 em 0.0000 f1 0.8000 rouge 0.4000\n", ""},
	{[]string{"score"}, "ref q1 aa\nref q2 bb\nans q1 m aa\nans q2 m cc\nans q2 m bb\nans q2 n bb\n", 0, "m samples 3 cases 2 em 0.6667 f1 0.6667 rouge 0.6667\nn samples 1 cases 1 em 1.0000 f1 1.0000 rouge 1.0000\n", ""},
	{[]string{"score"}, "ref q1 x\nans q1 Zed x\nans q1 mini x\nans q1 Abe y\nans q1 a1 x\n", 0, "Abe samples 1 cases 1 em 0.0000 f1 0.0000 rouge 0.0000\nZed samples 1 cases 1 em 1.0000 f1 1.0000 rouge 1.0000\na1 samples 1 cases 1 em 1.0000 f1 1.0000 rouge 1.0000\nmini samples 1 cases 1 em 1.0000 f1 1.0000 rouge 1.0000\n", ""},
	{[]string{"score"}, "ref q1 aa bb\nans q1 m aa\nans q1 m bb\nans q1 m bb aa\n", 0, "m samples 3 cases 1 em 0.0000 f1 0.7778 rouge 0.6111\n", ""},
	{[]string{"score"}, "ref q1 aa\nans q1 m aa\nans q1 m aa\nans q1 m bb\n", 0, "m samples 3 cases 1 em 0.6667 f1 0.6667 rouge 0.6667\n", ""},
	{[]string{"score"}, "", 0, "", ""},
	{[]string{"score"}, "ref q1 aa\nref q2 bb\n", 0, "", ""},
	{[]string{"score", "--metric", "rouge"}, "ref q1 aa bb\nans q1 m bb aa\nans q1 n aa bb\n", 0, "m samples 1 cases 1 rouge 0.5000\nn samples 1 cases 1 rouge 1.0000\n", ""},
}

func TestRowsTable(t *testing.T) { checkRows(t, rowsTable) }
