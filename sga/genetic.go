package sga

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MUTATIONRATE   = 0.1
	CROSSOVERRATE  = 0.8
	MAXGENERATIONS = 40
	TOURNAMENTSIZE = 3
)

const (
	MINIMIZE = iota
	MAXIMIZE
)

type obejectiveFunc func(float64) float64

type Genetic struct {
	minXvalue      float64
	maxXvalue      float64
	pool           population
	initialSize    int
	generation     int
	MutationRate   float64
	CrossoverRate  float64
	MaxGenerations int
	TournamentSize int
	target         float64
	Result         float64
	stats          Statistics
}

func Initialize(min, max float64, size int) (*Genetic, error) {
	if size <= 0 {
		return nil, fmt.Errorf("size cannot be less than or equal zero")
	}
	ga := new(Genetic)
	ga.minXvalue = min
	ga.maxXvalue = max
	ga.initialSize = size
	ga.generation = 1
	ga.pool.randomPopulation(min, max, size)
	return ga, nil
}

func (ga *Genetic) Maximize(objective obejectiveFunc) (float64, error) {
	ga.target = -1
	result, err := ga.solve(objective)
	ga.Result = result
	return result, err
}

func (ga *Genetic) Minimize(objective obejectiveFunc) (float64, error) {
	ga.target = 1
	return ga.solve(objective)
}

func (ga *Genetic) solve(objectiveFn obejectiveFunc) (float64, error) {
	if objectiveFn == nil {
		return 0, fmt.Errorf("null objective function")
	}
	if ga.target == 0 {
		return 0, fmt.Errorf("invalid target: the value must be 1 or -1")
	}
	if ga.MutationRate <= 0 {
		ga.MutationRate = MUTATIONRATE
	}
	if ga.CrossoverRate <= 0 {
		ga.CrossoverRate = CROSSOVERRATE
	}
	if ga.MaxGenerations <= 0 {
		ga.MaxGenerations = MAXGENERATIONS
	}
	if ga.TournamentSize <= 0 || ga.TournamentSize > ga.pool.size() {
		ga.TournamentSize = TOURNAMENTSIZE
	}
	//target represents the objective: maximize or minimize.
	//Any function can be minimized or maximized when multiplied by 1 or -1 respectively
	for i := 0; i < ga.MaxGenerations; i++ {
		ga.generation++
		ga.pool.fitness(objectiveFn, ga.target)
		ga.selection()
		ga.crossover()
	}
	ga.pool.fitness(objectiveFn, ga.target)
	return ga.pool.bestChromossome().chrom, nil
}

func (ga *Genetic) selection() error {
	matingPool := make(population, 0, len(ga.pool))
	//Elitism selection. The best chromosome is bring to next generationn
	best := ga.pool.bestChromossome()
	ga.stats.appendBest(best)
	matingPool.append(best)
	for i := 0; i < int((ga.pool.size()*2)/3)-1; i++ {
		best, err := ga.pool.tournament(ga.TournamentSize)
		if err != nil {
			return fmt.Errorf("fail to execute selection: %v", err)
		}
		matingPool.append(best)
	}
	ga.pool = matingPool
	return nil
}

func (ga *Genetic) crossover() {
	blx := func(p1, p2, alpha float64) float64 {
		rand.Seed(time.Now().UnixNano())
		beta := unifRandFloat(-1*alpha, alpha)
		return p1 + beta*(p2-p1)
	}
	size := ga.initialSize * 2 / 3
	for i := 0; i < size-1; i++ {
		var chield individual
		rand.Seed(time.Now().UnixNano())
		if rand.Float64() < ga.CrossoverRate {
			newChrom := blx(ga.pool[i].chrom, ga.pool[i+1].chrom, 0.5)
			for newChrom > ga.maxXvalue || newChrom < ga.minXvalue {
				newChrom = blx(ga.pool[i].chrom, ga.pool[i+1].chrom, 0.5)
			}
			chield = individual{chrom: newChrom}
			chield.mutate(ga.minXvalue, ga.maxXvalue, ga.MutationRate)
			ga.pool.append(chield)
		}
	}
}

func (ga *Genetic) Statistics() Statistics {
	return ga.stats
}
