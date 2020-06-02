package main

import (
	"fmt"
)

func fuelGuage(fuel int) {
	fmt.Printf("You have %d gallons of fuel.\n", fuel)
}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Printf("Greetings passengers! Today we are flying to %s\n", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.\n")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {

	fuel := 1000000
	planet := "Venus"
	fuel = flyToPlanet(planet, fuel)
	fuelGuage(fuel)
	planet = "Mars"
	fuel = flyToPlanet(planet, fuel)
	fuelGuage(fuel)
}
