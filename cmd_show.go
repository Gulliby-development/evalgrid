package main

import (
	"fmt"
	"strings"
)

func cmdShow(inv invocation) (outcome, *failure) {
	id := inv.args[0]
	reference, exists := inv.ledger.refs[id]
	if !exists {
		return outcome{}, fail(1, "no case "+id)
	}

	var output strings.Builder
	fmt.Fprintf(&output, "ref %s\n", reference)
	refTokens := tokens(reference)
	for _, sample := range inv.ledger.answers {
		if sample.id != id {
			continue
		}
		answerTokens := tokens(sample.text)
		match := 0
		if em(answerTokens, refTokens).Num().Sign() != 0 {
			match = 1
		}
		fmt.Fprintf(&output, "%s em %d f1 %s rouge %s\n",
			sample.model,
			match,
			format4(f1(answerTokens, refTokens)),
			format4(rouge(answerTokens, refTokens)))
	}
	return outcome{status: 0, out: output.String()}, nil
}
