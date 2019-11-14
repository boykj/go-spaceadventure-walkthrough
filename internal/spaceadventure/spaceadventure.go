package spaceadventure

import "fmt"

func Start(planetarySystem PlanetarySystem) {
	printWelcome(planetarySystem)
	printGreeting(responseToPrompt("Please Enter your name"))
	fmt.Println("Lets adventure")
	travel(promptForRandomOrSpecificDestination("Would you like to randomly travel to a planet?"))
}

func printWelcome(planetarySystem PlanetarySystem) {
	fmt.Println("Welcome to the %s!\n", planetarySystem.Name)
	fmt.Println("There are %d planets to explore.\n", planetarySystem.NumberOfPlanets())
}

func getName() string{
	var name string
	fmt.Println("What is your name?")
    fmt.Scan(&name)
    return name
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel(randomDestination bool) {
	if (randomDestination) {
		travelToRandomPlanet();
	} else {
		travelToPlanet(responseToPrompt("Name the planet you woud like to travel to"))
	}
}

func promptForRandomOrSpecificDestination(prompt string) bool {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = responseToPrompt(prompt)
		if choice == "Y" {
			return true
		} else if choice == "N" {
			return false
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
	return false
}

func responseToPrompt(prompt string) string{
	var response string
	fmt.Println(prompt)
	fmt.Scan(&response)
	return response
}

func travelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}
