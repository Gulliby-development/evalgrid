package main

import "math/big"

func em(p, r []string) *big.Rat {
	if len(p) != len(r) {
		return new(big.Rat)
	}
	for i := range p {
		if p[i] != r[i] {
			return new(big.Rat)
		}
	}
	return new(big.Rat).SetInt64(1)
}

func f1(p, r []string) *big.Rat {
	if len(p) == 0 && len(r) == 0 {
		return new(big.Rat).SetInt64(1)
	}
	if len(p) == 0 || len(r) == 0 {
		return new(big.Rat)
	}

	counts := make(map[string]int, len(p))
	for _, token := range p {
		counts[token]++
	}
	overlap := 0
	for _, token := range r {
		if counts[token] > 0 {
			overlap++
			counts[token]--
		}
	}
	return overlapScore(overlap, len(p)+len(r))
}

func rouge(p, r []string) *big.Rat {
	if len(p) == 0 && len(r) == 0 {
		return new(big.Rat).SetInt64(1)
	}
	if len(p) == 0 || len(r) == 0 {
		return new(big.Rat)
	}

	// The row represents the preceding prefix of r; each update computes
	// the LCS length for the corresponding prefixes of p and r.
	previous := make([]int, len(r)+1)
	for _, pToken := range p {
		current := make([]int, len(r)+1)
		for j, rToken := range r {
			if pToken == rToken {
				current[j+1] = previous[j] + 1
			} else if previous[j+1] > current[j] {
				current[j+1] = previous[j+1]
			} else {
				current[j+1] = current[j]
			}
		}
		previous = current
	}
	return overlapScore(previous[len(r)], len(p)+len(r))
}

func overlapScore(overlap, total int) *big.Rat {
	return new(big.Rat).SetFrac(big.NewInt(int64(2*overlap)), big.NewInt(int64(total)))
}

func mean(values []*big.Rat) *big.Rat {
	total := new(big.Rat)
	for _, value := range values {
		total.Add(total, value)
	}
	if len(values) == 0 {
		return total
	}
	return total.Quo(total, new(big.Rat).SetInt64(int64(len(values))))
}
