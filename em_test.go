package main

import "testing"

var rowsEm = []row{
	{[]string{"score", "--metric", "em"}, "ref q1 yes\nans q1 m yes\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 yes\nans q1 m no\n", 0, "m samples 1 cases 1 em 0.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 yes\nans q1 m yes\nans q1 m no\nans q1 m yes\n", 0, "m samples 3 cases 1 em 0.6667\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 a1\nref q2 b2\nans q1 m a1\nans q2 m x\nans q2 m b2\nans q2 m b2 b2\n", 0, "m samples 4 cases 2 em 0.5000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 a1\nans q1 m a1\nans q1 n a1\nans q1 m b\n", 0, "m samples 2 cases 1 em 0.5000\nn samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 x\nans q1 b x\nans q1 B x\nans q1 a x\nans q1 _ x\nans q1 a.b x\n", 0, "B samples 1 cases 1 em 1.0000\n_ samples 1 cases 1 em 1.0000\na samples 1 cases 1 em 1.0000\na.b samples 1 cases 1 em 1.0000\nb samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 x\nref q2 y\n", 0, "", ""},
	{[]string{"score", "--metric", "em"}, "", 0, "", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 x\nans q1 m y\nans q1 m y\nans q1 m x\n", 0, "m samples 3 cases 1 em 0.3333\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 one two\nans q1 m two one\nans q1 m one two two\n", 0, "m samples 2 cases 1 em 0.0000\n", ""},
}

func TestRowsEm(t *testing.T) { checkRows(t, rowsEm) }
