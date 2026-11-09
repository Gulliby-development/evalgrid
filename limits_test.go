package main

import "testing"

var rowsLimits = []row{
	{[]string{"score"}, "ref abcdefghijklmnopqrst x\nans abcdefghijklmnopqrst m x\n", 0, "m samples 1 cases 1 em 1.0000 f1 1.0000 rouge 1.0000\n", ""},
	{[]string{"score"}, "ref abcdefghijklmnopqrstu x\nans q1 m x\n", 1, "", "evalgrid: line 1: bad id abcdefghijklmnopqrstu\n"},
	{[]string{"score"}, "ref q1 x\nans abcdefghijklmnopqrstu m x\n", 1, "", "evalgrid: line 2: bad id abcdefghijklmnopqrstu\n"},
	{[]string{"score"}, "ref q1 x\nans q1 abcdefghijklmnopqrstuvwxyzABCD0123456789 x\n", 0, "abcdefghijklmnopqrstuvwxyzABCD0123456789 samples 1 cases 1 em 1.0000 f1 1.0000 rouge 1.0000\n", ""},
	{[]string{"score"}, "ref q1 x\nans q1 abcdefghijklmnopqrstuvwxyzABCD0123456789x x\n", 1, "", "evalgrid: line 2: bad model name abcdefghijklmnopqrstuvwxyzABCD0123456789x\n"},
	{[]string{"score", "--metric", "em"}, "ref q1 abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 012345\nans q1 m abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 012345\n", 0, "m samples 1 cases 1 em 1.0000\n", ""},
	{[]string{"score"}, "ref q1 abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456\n", 1, "", "evalgrid: line 1: text too long\n"},
	{[]string{"score"}, "ref q1 x\nans q1 m   abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 012345   \n", 0, "m samples 1 cases 1 em 0.0000 f1 0.0000 rouge 0.0000\n", ""},
	{[]string{"score", "--metric", "f1"}, "ref q1 x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\nans q1 m x\n", 0, "m samples 40 cases 1 f1 1.0000\n", ""},
}

func TestRowsLimits(t *testing.T) { checkRows(t, rowsLimits) }
