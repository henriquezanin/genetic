# Genetic

Genetic is a Go library for function optimization.

## Installation

Use the Go package manager to install Genetic.

```bash
go get github.com/henriquezanin/genetic
```

## Usage

```go
import (
    ...
    genetic "github.com/henriquezanin/genetic/sga"
)
func main() {
    //Min x value, max x value, population size
    ga, err := genetic.Initialize(-1, 2, 30)
    if err != nil {
        log.Fatalf("Failed to initalize population: %v", err)
    }
    //Definition of objetive function
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
```

## License
[GPLv3](https://choosealicense.com/licenses/gpl-3.0/)