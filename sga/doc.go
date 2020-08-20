// Copyright 2020 Henrique Gomes Zanin. All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

/*
Package genetic provides a optimization framework using Genetic Algorithms

Initialize, Maximize, Minimize and Statistics are operations to perform a optimizations.

	ga, err = genetic.Initialize(-1, 2, 30)
	...
	result, err = population.Maximize(objective function)
	...
	result, err = population.Minimize(objective function)
	...
	statistics = ga.Statistics()

The Initialize function needs the parameters: minimum x value, maximum x value and population size. As return gives a type Genetic thats contain Maximize, Minimize and Statistics methods.
Maximize and Minimize method needs an objective function defined by user.
Both methods, Maximize and Minimize, returns a real number as optimization result and error. On the other hand Statistics method provide an type Statistics that allows user to
plot the convergence process and print statistics data.

Genetic algorithms parameters:

	int value to define maximum iteration number
		ga.MaxGenerations

	float64 value to change mutation rate thats represents the percentage of children is mutate. Must be between 0 and 1
		ga.MutationRate

	float64 thats allow changes at crossover rate
		ga.CrossoverRate

	The tournament size modify the selection pression. Must be between 1 and population size
		ga.TournamentSize

*/
package sga
