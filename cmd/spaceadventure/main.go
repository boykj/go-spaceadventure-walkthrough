package main

import "github.com/boykj/go-spaceadventure-walkthrough/internal/spaceadventure"


func main() {
	// spaceadventure.Start(spaceadventure.Planetarysystem{"Salad system"}) 
	// var ps = spaceadventure.PlanetarySystem{"Salad system"}

	// Naming can also be done this way
	// var ps = spaceadventure.PlanetarySystem{"Salad system"}
	// spaceadventure.Start(ps)
	
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name: "Salad System", Planets: []spaceadventure.Planet{
				spaceadventure.Planet{"Tatooine", "Desert Planet"}
			},
		},
	)

}
