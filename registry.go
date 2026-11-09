package main

const usage = "usage: evalgrid score [--metric NAME]\n" +
	"       evalgrid passk K\n" +
	"       evalgrid show ID\n" +
	"       evalgrid check [--min PCT]\n"

// invocation is what a command gets: the value of its option ("" when not given),
// its positional arguments and the parsed ledger.
type invocation struct {
	option string
	args   []string
	ledger ledger
}

// outcome is a command's result: the exit status and the text for standard output.
type outcome struct {
	status int
	out    string
}

type command struct {
	name       string
	option     string
	positional int
	needs      string
	run        func(inv invocation) (outcome, *failure)
}

var commands = []command{
	{"score", "--metric", 0, "", cmdScore},
	{"passk", "", 1, "K", cmdPassk},
	{"show", "", 1, "ID", cmdShow},
	{"check", "--min", 0, "", cmdCheck},
}

func findCommand(name string) *command {
	for i := range commands {
		if commands[i].name == name {
			return &commands[i]
		}
	}
	return nil
}
