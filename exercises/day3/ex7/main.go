package main

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Insurable interface {
	CalculateInsurance() float64
}

type Printable interface {
	Details()
}

type Car struct {
	Vehicle
	NumberOfDoors int
}

func (c Car) CalculateInsurance() float64 {
	return 1000
}

func (c Car) Details() {
	fmt.Printf("Car = Make: %s, Model: %s, Year: %d, Doors: %d\n", c.Make, c.Model, c.Year, c.NumberOfDoors)
}

type Truck struct {
	Vehicle
	PayloadCapacity int
}

func (t Truck) CalculateInsurance() float64 {
	return 5000
}

func (t Truck) Details() {
	fmt.Printf("Truck = Make: %s, Model: %s, Year: %d, Payload Capacity: %d\n", t.Make, t.Model, t.Year, t.PayloadCapacity)
}

func PrintAll(p []Printable) {
	for _, pr := range p {
		pr.Details()
	}
}
func main() {
	car := Car{
		Vehicle: Vehicle{
			Make:  "Toyota",
			Model: "Corolla",
			Year:  2020,
		},
		NumberOfDoors: 4,
	}

	truck := Truck{
		Vehicle: Vehicle{
			Make:  "Ford",
			Model: "F-150",
			Year:  2018,
		},
		PayloadCapacity: 3000,
	}

	objs := []Printable{car, truck}

	PrintAll(objs)
}
