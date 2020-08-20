package sga

import (
	"fmt"
	"math/rand"
	"time"
)

func unifRandFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

func (i individual) String() string {
	return fmt.Sprintf("Chromossome: %.4f,\tObjective Value: %.4f\tFitness: %.4f", i.chrom, i.objectiveValue, i.fitness)
}

func (p *Genetic) String() string {
	info := fmt.Sprintf("Generation: %d\n", p.generation)
	info = fmt.Sprintf("%vGenetic:", info)
	for i, each := range p.pool {
		info = fmt.Sprintf("%v\n[%d]\t%v", info, i+1, each)
	}
	return info
}
