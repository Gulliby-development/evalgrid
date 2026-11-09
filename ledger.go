package main

import (
	"strconv"
	"strings"
)

type ledger struct {
	refs    map[string]string
	answers []answer
}

type answer struct {
	id    string
	model string
	text  string
	line  int
	order int
}

// validID accepts the lowercase ASCII identifier alphabet used by the ledger.
func validID(text string) bool {
	if len(text) < 1 || len(text) > 20 {
		return false
	}
	for i := 0; i < len(text); i++ {
		c := text[i]
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' || c == '-') {
			return false
		}
	}
	return true
}

func validModelName(text string) bool {
	if len(text) < 1 || len(text) > 40 {
		return false
	}
	for i := 0; i < len(text); i++ {
		c := text[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || strings.ContainsRune("._:/-", rune(c))) {
			return false
		}
	}
	return true
}

type ledgerWord struct {
	value string
	end   int
}

func lineWords(line string) []ledgerWord {
	words := make([]ledgerWord, 0)
	for i := 0; i < len(line); {
		for i < len(line) && (line[i] == ' ' || line[i] == '\t') {
			i++
		}
		if i == len(line) {
			break
		}
		start := i
		for i < len(line) && line[i] != ' ' && line[i] != '\t' {
			i++
		}
		words = append(words, ledgerWord{value: line[start:i], end: i})
	}
	return words
}

func trimBlanks(text string) string {
	start, end := 0, len(text)
	for start < end && (text[start] == ' ' || text[start] == '\t') {
		start++
	}
	for end > start && (text[end-1] == ' ' || text[end-1] == '\t') {
		end--
	}
	return text[start:end]
}

func badLedgerChar(line string) bool {
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c != '\t' && (c < 32 || c > 126) {
			return true
		}
	}
	return false
}

func parseLedger(text string) (ledger, *failure) {
	result := ledger{refs: make(map[string]string)}
	lines := strings.Split(text, "\n")
	if strings.HasSuffix(text, "\n") {
		lines = lines[:len(lines)-1]
	}
	for i, line := range lines {
		lineNumber := i + 1
		if strings.HasSuffix(text, "\n") || i < len(lines)-1 {
			if strings.HasSuffix(line, "\r") {
				line = line[:len(line)-1]
			}
		}
		if badLedgerChar(line) {
			return ledger{}, fail(1, lineError(lineNumber, "bad character"))
		}
		words := lineWords(line)
		if len(words) == 0 || strings.HasPrefix(words[0].value, "#") {
			continue
		}
		kind := words[0].value
		if kind != "ref" && kind != "ans" {
			return ledger{}, fail(1, lineError(lineNumber, "unknown line kind "+kind))
		}
		if kind == "ref" {
			if len(words) < 2 {
				return ledger{}, fail(1, lineError(lineNumber, "ref needs an id"))
			}
			id := words[1].value
			if !validID(id) {
				return ledger{}, fail(1, lineError(lineNumber, "bad id "+id))
			}
			answerText := trimBlanks(line[words[1].end:])
			if answerText == "" {
				return ledger{}, fail(1, lineError(lineNumber, "ref "+id+" needs text"))
			}
			if len(answerText) > 60 {
				return ledger{}, fail(1, lineError(lineNumber, "text too long"))
			}
			if _, exists := result.refs[id]; exists {
				return ledger{}, fail(1, lineError(lineNumber, "ref for "+id+" defined twice"))
			}
			result.refs[id] = answerText
			continue
		}
		if len(words) < 2 {
			return ledger{}, fail(1, lineError(lineNumber, "ans needs an id"))
		}
		id := words[1].value
		if !validID(id) {
			return ledger{}, fail(1, lineError(lineNumber, "bad id "+id))
		}
		if len(words) < 3 {
			return ledger{}, fail(1, lineError(lineNumber, "ans needs a model name"))
		}
		model := words[2].value
		if !validModelName(model) {
			return ledger{}, fail(1, lineError(lineNumber, "bad model name "+model))
		}
		answerText := trimBlanks(line[words[2].end:])
		if len(answerText) > 60 {
			return ledger{}, fail(1, lineError(lineNumber, "text too long"))
		}
		result.answers = append(result.answers, answer{id: id, model: model, text: answerText, line: lineNumber, order: len(result.answers)})
	}
	for _, a := range result.answers {
		if _, exists := result.refs[a.id]; !exists {
			return ledger{}, fail(1, lineError(a.line, "no reference for "+a.id))
		}
	}
	return result, nil
}

func lineError(line int, message string) string {
	return "line " + strconv.Itoa(line) + ": " + message
}
