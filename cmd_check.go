package main

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

func cmdCheck(inv invocation) (outcome, *failure) {
	minimum := int64(50)
	if inv.option != "" {
		// CLI validation guarantees this is a decimal integer from 0 through 100.
		minimum, _ = strconv.ParseInt(inv.option, 10, 64)
	}

	scores := make(map[string][]*big.Rat)
	for _, sample := range inv.ledger.answers {
		scores[sample.model] = append(scores[sample.model], em(tokens(sample.text), tokens(inv.ledger.refs[sample.id])))
	}
	if len(scores) == 0 {
		return outcome{status: 0, out: "no answers\n"}, nil
	}

	models := make([]string, 0, len(scores))
	for model := range scores {
		models = append(models, model)
	}
	sort.Strings(models)

	threshold := new(big.Rat).SetInt64(minimum)
	var output strings.Builder
	below := false
	for _, model := range models {
		percentage := new(big.Rat).Mul(mean(scores[model]), new(big.Rat).SetInt64(100))
		status := "ok"
		if percentage.Cmp(threshold) < 0 {
			status = "below"
			below = true
		}
		fmt.Fprintf(&output, "%s em %s%% %s\n", model, format2(percentage), status)
	}
	code := 0
	if below {
		code = 3
	}
	return outcome{status: code, out: output.String()}, nil
}
