package main

import (
	"encoding/json"
	"log"
	"math"
	"os"
)

type Person struct {
	Name      string
	Age       int
	Salary    float64
	Education string
}
type People struct {
	Person []Person
}

type Stats struct {
	AverageAge            float64
	YoungestPersons       []Person
	OldestPersons         []Person
	AverageSalary         float64
	HighestPaidPersons    []Person
	LowestPaidPersons     []Person
	CountByEducationLevel map[string]int
}
type StatsWrapper struct {
	Stats []Stats
}

func generateStats(people []Person) Stats {

	avAge := func() float64 {
		var sum float64
		for _, p := range people {
			sum += float64(p.Age)
		}
		return sum / float64(len(people))
	}

	youngest := func() []Person {
		mn := math.MaxInt
		yngs := []Person{}
		for _, p := range people {
			mn = min(mn, p.Age)
		}
		for _, p := range people {
			if p.Age == mn {
				yngs = append(yngs, p)
			}
		}
		return yngs
	}

	oldest := func() []Person {
		mx := -1
		olds := []Person{}
		for _, p := range people {
			mx = max(mx, p.Age)
		}
		for _, p := range people {
			if p.Age == mx {
				olds = append(olds, p)
			}
		}
		return olds
	}
	avSalary := func() float64 {
		var sum float64
		for _, p := range people {
			sum += float64(p.Salary)
		}
		return sum / float64(len(people))
	}

	maxSalary := func() []Person {
		var mx float64 = -1
		mxs := []Person{}
		for _, p := range people {
			mx = max(mx, p.Salary)
		}
		for _, p := range people {
			if p.Salary == mx {
				mxs = append(mxs, p)
			}
		}
		return mxs
	}
	minSalary := func() []Person {
		mn := math.MaxFloat64
		mns := []Person{}
		for _, p := range people {
			mn = min(mn, p.Salary)
		}
		for _, p := range people {
			if p.Salary == mn {
				mns = append(mns, p)
			}
		}
		return mns
	}

	countByEd := func() map[string]int {
		freq := map[string]int{}

		for _, p := range people {
			_, f := freq[p.Education]
			if f {
				freq[p.Education]++
			} else {
				freq[p.Education] = 1
			}
		}
		return freq

	}

	return Stats{avAge(), youngest(), oldest(), avSalary(), maxSalary(), minSalary(), countByEd()}
}
func main() {
	file, err := os.ReadFile("./person.json")
	if err != nil {
		log.Fatal(err)
	}

	people := People{}
	err = json.Unmarshal(file, &people)
	if err != nil {
		log.Fatal(err)
	}
	wrapper := StatsWrapper{}
	wrapper.Stats = append(wrapper.Stats, generateStats(people.Person))

	pa, err := json.MarshalIndent(wrapper, "  ", "	  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./stats.json", pa, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
