package main

import "testing"

var rowsLines = []row{
	{[]string{"score"}, "rate q1 x\n", 1, "", "evalgrid: line 1: unknown line kind rate\n"},
	{[]string{"score"}, "# note\n\nREF q1 x\n", 1, "", "evalgrid: line 3: unknown line kind REF\n"},
	{[]string{"score"}, "ref\n", 1, "", "evalgrid: line 1: ref needs an id\n"},
	{[]string{"score"}, "ref Q1 x\n", 1, "", "evalgrid: line 1: bad id Q1\n"},
	{[]string{"score"}, "ref q1\n", 1, "", "evalgrid: line 1: ref q1 needs text\n"},
	{[]string{"score"}, "ref q1   \t \n", 1, "", "evalgrid: line 1: ref q1 needs text\n"},
	{[]string{"score"}, "ans\n", 1, "", "evalgrid: line 1: ans needs an id\n"},
	{[]string{"score"}, "ans q1\n", 1, "", "evalgrid: line 1: ans needs a model name\n"},
	{[]string{"score"}, "ref q1 x\nans q1 bad@model y\n", 1, "", "evalgrid: line 2: bad model name bad@model\n"},
	{[]string{"score"}, "ref q1 x\nref q1 y\n", 1, "", "evalgrid: line 2: ref for q1 defined twice\n"},
	{[]string{"score"}, "ref q1 x\nans q2 m y\n", 1, "", "evalgrid: line 2: no reference for q2\n"},
	{[]string{"score"}, "ans q2 m y\nref q2 x\nans q3 m z\n", 1, "", "evalgrid: line 3: no reference for q3\n"},
	{[]string{"score"}, "ref q1 x\n# c\n\n  \nans q1 m y\nans q9 m y\nbad\n", 1, "", "evalgrid: line 7: unknown line kind bad\n"},
	{[]string{"score"}, "ans q1 m y\nref q1 x\n", 0, "m samples 1 cases 1 em 0.0000 f1 0.0000 rouge 0.0000\n", ""},
}

func TestRowsLines(t *testing.T) { checkRows(t, rowsLines) }
