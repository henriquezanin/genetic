package sga

import "fmt"

type Statistics struct {
	bestOfGenerations population
}

func (stat Statistics) NumberOfGenerations() int {
	return stat.bestOfGenerations.size()
}

func (stat *Statistics) appendBest(ind individual) {
	stat.bestOfGenerations.append(ind)
}

func (stat Statistics) ShowConvergence() {
	for i, each := range stat.bestOfGenerations {
		fmt.Printf("[%d] - %f\n", i, each.chrom)
	}
}
