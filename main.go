package main

import (
	"fmt"
	"log"
	"math"

	genetic "github.com/henriquezanin/genetic/sga"
)

func main() {
	ga, err := genetic.Initialize(-1, 2, 30)
	if err != nil {
		log.Fatalf("Failed to initalize population: %v", err)
	}
	obj := func(x float64) float64 {
		return x*math.Sin(10*math.Pi*x) + 1
	}
	ga.MaxGenerations = 200
	ga.MutationRate = 0.1
	ga.CrossoverRate = 0.7
	ga.TournamentSize = 3
	ga.Maximize(obj)
	fmt.Println(ga.Result)
	stats := ga.Statistics()
	stats.ShowConvergence()
}
