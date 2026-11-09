package main

import "testing"

var rowsNorm = []row{
	{[]string{"score", "--metric", "em"}, "ref q1 Paris\nans q1 m paris\nans q1 m PARIS\n", 0, "m samples 2 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 the cat\nans q1 m cat\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 cat\nans q1 m A cat\nans q1 m an cat\nans q1 m the cat\nans q1 m thecat\nans q1 m another cat\n", 0, "m samples 5 cases 1 em 0.6000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 Hello, world!\nans q1 m hello world\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 well-known\nans q1 m wellknown\nans q1 m well known\n", 0, "m samples 2 cases 1 em 0.5000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 it's\nans q1 m its\nans q1 m it s\n", 0, "m samples 2 cases 1 em 0.5000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 the\nans q1 m\nans q1 m a an the\n", 0, "m samples 2 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 the\nans q1 m\nans q1 m cat\n", 0, "m samples 2 cases 1 f1 0.5000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 (a) [b] {c}\nans q1 m b c\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 $5.00 + 3\nans q1 m 500 3\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score", "--metric", "em"}, "ref q1 x_y a.b\nans q1 m xy ab\nans q1 m x y a b\n", 0, "m samples 2 cases 1 em 0.5000\n", ""},
}

func TestRowsNorm(t *testing.T) { checkRows(t, rowsNorm) }
