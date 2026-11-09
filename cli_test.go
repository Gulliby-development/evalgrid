package main

import "testing"

var rowsCli = []row{
	{nil, "", 2, "", "usage: evalgrid score [--metric NAME]\n       evalgrid passk K\n       evalgrid show ID\n       evalgrid check [--min PCT]\n"},
	{[]string{"help"}, "", 0, "usage: evalgrid score [--metric NAME]\n       evalgrid passk K\n       evalgrid show ID\n       evalgrid check [--min PCT]\n", ""},
	{[]string{"-h"}, "", 0, "usage: evalgrid score [--metric NAME]\n       evalgrid passk K\n       evalgrid show ID\n       evalgrid check [--min PCT]\n", ""},
	{[]string{"--help"}, "", 0, "usage: evalgrid score [--metric NAME]\n       evalgrid passk K\n       evalgrid show ID\n       evalgrid check [--min PCT]\n", ""},
	{[]string{"help"}, "ref a\n", 0, "usage: evalgrid score [--metric NAME]\n       evalgrid passk K\n       evalgrid show ID\n       evalgrid check [--min PCT]\n", ""},
	{[]string{"--version"}, "", 0, "evalgrid 0.1.0\n", ""},
	{[]string{"--version", "x"}, "ref q1 the cat sat\nans q1 m cat sat\n", 0, "evalgrid 0.1.0\n", ""},
	{[]string{"frob"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: unknown command frob\n"},
	{[]string{""}, "", 2, "", "evalgrid: unknown command \n"},
	{[]string{"Score"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: unknown command Score\n"},
	{[]string{"-x"}, "", 2, "", "evalgrid: unknown command -x\n"},
	{[]string{"score", "--version"}, "ref q1 the cat sat\nans q1 m cat sat\n", 2, "", "evalgrid: unknown option --version\n"},
}

func TestRowsCli(t *testing.T) { checkRows(t, rowsCli) }
