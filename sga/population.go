package sga

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//Type individual represents an potential solution
type individual struct {
	chrom          float64
	fitness        float64
	objectiveValue float64
}

type population []individual

func (pop *population) randomPopulation(min, max float64, size int) {
	*pop = make(population, 0, size)
	for i := 0; i < size; i++ {
		temp := individual{
			chrom: unifRandFloat(min, max),
		}
		*pop = append(*pop, temp)
	}
}

func (pop *population) append(ind individual) {
	*pop = append(*pop, ind)
}

func (pop *population) fitness(objective obejectiveFunc, target float64) {
	for i, each := range *pop {
		//target define if the objective of GA is maximization or minimization
		(*pop)[i].objectiveValue = objective(each.chrom)
		(*pop)[i].fitness = target * (*pop)[i].objectiveValue
	}
}

func (pop population) tournament(size int) (individual, error) {
	var best individual
	if size > pop.size() {
		return best, fmt.Errorf("tournament size bigger than population size")
	}
	for i := 0; i < size; i++ {
		point := int(math.Round(unifRandFloat(0, float64(pop.size()-1))))
		if pop[point].fitness < best.fitness {
			best = pop[point]
		}
	}
	return best, nil
}

func (pop population) bestChromossome() individual {
	var best individual
	for _, each := range pop {
		if each.fitness < best.fitness {
			best = each
		}
	}
	return best
}

func (pop population) size() int {
	return len(pop)
}

func (ind *individual) mutate(lowerBound, upperBound, probability float64) float64 {
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < probability {
		ind.chrom = unifRandFloat(lowerBound, upperBound)
	}
	return 0
}
