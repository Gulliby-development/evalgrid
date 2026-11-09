package main

import "strconv"

func readsStdin(args []string) bool {
	if len(args) == 0 || isInformational(args[0]) {
		return false
	}
	cmd := findCommand(args[0])
	if cmd == nil {
		return false
	}
	_, _, f := validateCommand(cmd, args[1:])
	return f == nil
}

func run(args []string, input string) (int, string, string) {
	if len(args) == 0 {
		return 2, "", usage
	}
	if isInformational(args[0]) {
		if args[0] == "--version" {
			return 0, "evalgrid 0.1.0\n", ""
		}
		return 0, usage, ""
	}
	cmd := findCommand(args[0])
	if cmd == nil {
		return 2, "", "evalgrid: unknown command " + args[0] + "\n"
	}
	inv, _, f := validateCommand(cmd, args[1:])
	if f != nil {
		return f.code, "", "evalgrid: " + f.msg + "\n"
	}
	parsed, f := parseLedger(input)
	if f != nil {
		return f.code, "", "evalgrid: " + f.msg + "\n"
	}
	inv.ledger = parsed
	result, f := cmd.run(inv)
	if f != nil {
		return f.code, "", "evalgrid: " + f.msg + "\n"
	}
	return result.status, result.out, ""
}

func isInformational(word string) bool {
	return word == "help" || word == "-h" || word == "--help" || word == "--version"
}

// validateCommand returns the invocation, whether an option was supplied, and the first CLI failure.
func validateCommand(cmd *command, words []string) (invocation, bool, *failure) {
	inv := invocation{}
	optionSet := false
	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(word) > 1 && word[0] == '-' {
			if cmd.option == "" || word != cmd.option {
				return inv, false, fail(2, "unknown option "+word)
			}
			if optionSet {
				return inv, false, fail(2, word+" given twice")
			}
			optionSet = true
			if i+1 == len(words) || words[i+1] == "" {
				return inv, false, fail(2, word+" needs a value")
			}
			i++
			inv.option = words[i]
			continue
		}
		if len(inv.args) >= cmd.positional {
			return inv, false, fail(2, "unexpected argument "+word)
		}
		inv.args = append(inv.args, word)
	}
	if len(inv.args) < cmd.positional {
		return inv, optionSet, fail(2, cmd.name+" needs "+cmd.needs)
	}
	if optionSet {
		switch cmd.option {
		case "--metric":
			if inv.option != "em" && inv.option != "f1" && inv.option != "rouge" {
				return inv, optionSet, fail(2, "bad metric "+inv.option)
			}
		case "--min":
			if !asciiDigits(inv.option) || len(inv.option) > 3 {
				return inv, optionSet, fail(2, "--min needs an integer from 0 to 100")
			}
			value, _ := strconv.Atoi(inv.option)
			if value > 100 {
				return inv, optionSet, fail(2, "--min needs an integer from 0 to 100")
			}
		}
	}
	if cmd.name == "passk" {
		value := inv.args[0]
		if len(value) > 2 || !asciiDigits(value) {
			return inv, optionSet, fail(2, "bad k "+value)
		}
		n, _ := strconv.Atoi(value)
		if n < 1 || n > 99 {
			return inv, optionSet, fail(2, "bad k "+value)
		}
	}
	if cmd.name == "show" && !validID(inv.args[0]) {
		return inv, optionSet, fail(2, "bad id "+inv.args[0])
	}
	return inv, optionSet, nil
}

func asciiDigits(text string) bool {
	if text == "" {
		return false
	}
	for i := 0; i < len(text); i++ {
		if text[i] < '0' || text[i] > '9' {
			return false
		}
	}
	return true
}
