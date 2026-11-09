package main

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type passkGroup struct {
	samples   int
	successes int
}

func cmdPassk(inv invocation) (outcome, *failure) {
	k, _ := strconv.Atoi(inv.args[0])
	groups := make(map[string]map[string]*passkGroup)
	for _, sample := range inv.ledger.answers {
		modelGroups := groups[sample.model]
		if modelGroups == nil {
			modelGroups = make(map[string]*passkGroup)
			groups[sample.model] = modelGroups
		}
		group := modelGroups[sample.id]
		if group == nil {
			group = &passkGroup{}
			modelGroups[sample.id] = group
		}
		group.samples++
		if em(tokens(sample.text), tokens(inv.ledger.refs[sample.id])).Cmp(new(big.Rat).SetInt64(1)) == 0 {
			group.successes++
		}
	}

	models := make([]string, 0, len(groups))
	for model := range groups {
		models = append(models, model)
	}
	sort.Strings(models)

	var output strings.Builder
	for _, model := range models {
		var estimates []*big.Rat
		skipped := 0
		for _, group := range groups[model] {
			if group.samples < k {
				skipped++
				continue
			}
			estimates = append(estimates, passkEstimate(group.samples, group.successes, k))
		}
		value := "-"
		if len(estimates) > 0 {
			value = format4(mean(estimates))
		}
		fmt.Fprintf(&output, "%s cases %d skipped %d pass@%d %s\n",
			model, len(estimates), skipped, k, value)
	}
	return outcome{status: 0, out: output.String()}, nil
}

func passkEstimate(samples, successes, k int) *big.Rat {
	denominator := new(big.Int).Binomial(int64(samples), int64(k))
	numerator := new(big.Int)
	if samples-successes >= k {
		numerator.Binomial(int64(samples-successes), int64(k))
	}
	failureProbability := new(big.Rat).SetFrac(numerator, denominator)
	return new(big.Rat).Sub(new(big.Rat).SetInt64(1), failureProbability)
}
