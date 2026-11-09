package main

import "testing"

var rowsF1 = []row{
	{[]string{"score", "--metric", "f1"}, "ref q1 cat sat\nans q1 m cat sat\n", 0, "m samples 1 cases 1 f1 1.0000\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 cat sat\nans q1 m cat\n", 0, "m samples 1 cases 1 f1 0.6667\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 cat sat\nans q1 m dog ran\n", 0, "m samples 1 cases 1 f1 0.0000\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 cat sat on mat\nans q1 m the cat sat\n", 0, "m samples 1 cases 1 f1 0.6667\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 cat cat dog\nans q1 m cat dog dog\n", 0, "m samples 1 cases 1 f1 0.6667\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 cat\nans q1 m cat cat cat\n", 0, "m samples 1 cases 1 f1 0.5000\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 a b c d e f\nans q1 m b d f z\n", 0, "m samples 1 cases 1 f1 0.6667\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 cat sat\nans q1 m sat cat\n", 0, "m samples 1 cases 1 f1 1.0000\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 one two three\nans q1 m one\nans q1 m one two\nans q1 m\n", 0, "m samples 3 cases 1 f1 0.4333\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 b c d e f g h i j k l m n o p q\nans q1 m b r s t u v w x y z 0 1 2 3 4 5\nans q1 m zz\n", 0, "m samples 2 cases 1 f1 0.0313\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 aa bb cc\nans q1 m aa\nans q1 m bb\nans q1 m cc\n", 0, "m samples 3 cases 1 f1 0.5000\n", ""},
}

func TestRowsF1(t *testing.T) { checkRows(t, rowsF1) }
