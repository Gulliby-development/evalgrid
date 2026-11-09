package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

type scoreGroup struct {
	cases map[string]struct{}
	em    []*big.Rat
	f1    []*big.Rat
	rouge []*big.Rat
}

func cmdScore(inv invocation) (outcome, *failure) {
	groups := make(map[string]*scoreGroup)
	for _, sample := range inv.ledger.answers {
		group := groups[sample.model]
		if group == nil {
			group = &scoreGroup{cases: make(map[string]struct{})}
			groups[sample.model] = group
		}
		group.cases[sample.id] = struct{}{}
		p := tokens(sample.text)
		r := tokens(inv.ledger.refs[sample.id])
		group.em = append(group.em, em(p, r))
		group.f1 = append(group.f1, f1(p, r))
		group.rouge = append(group.rouge, rouge(p, r))
	}

	models := make([]string, 0, len(groups))
	for model := range groups {
		models = append(models, model)
	}
	sort.Strings(models)

	var output strings.Builder
	for _, model := range models {
		group := groups[model]
		fmt.Fprintf(&output, "%s samples %d cases %d", model, len(group.em), len(group.cases))
		if inv.option == "" {
			fmt.Fprintf(&output, " em %s f1 %s rouge %s\n",
				format4(mean(group.em)),
				format4(mean(group.f1)),
				format4(mean(group.rouge)))
			continue
		}
		var values []*big.Rat
		switch inv.option {
		case "em":
			values = group.em
		case "f1":
			values = group.f1
		case "rouge":
			values = group.rouge
		}
		fmt.Fprintf(&output, " %s %s\n", inv.option, format4(mean(values)))
	}
	return outcome{status: 0, out: output.String()}, nil
}
