// Package main is the entry point of the Go program.
package main

import (
	"fmt"
	"time"
)

// The main function is the entry point of the program.
func main() {
	// BUILT-IN TYPES
	// Declare and initialize variables using the var keyword.
	var ninja = "Johnny"
	var level, yoe int = 1, 2
	var isSkilled bool
	weapon := "Ninja Star"
	fmt.Println(ninja, level, yoe, isSkilled, weapon)

	// CONSTANTS
	// Declare constants using the const keyword.
	const dojo string = "Golang Dojo"
	const powerLevel = 9001
	const opLevel = 3e20
	fmt.Printf("%T\n", opLevel)

	// LOOPS
	// Use for loops to iterate over a range of values.
	isSkilled = true
	for isSkilled {
		fmt.Println("Ready for mission!")
		isSkilled = false
	}
	for level := 7; level < 9; level++ {
		fmt.Println(level)
		fmt.Println("Leveling up!")
	}
	for {
		fmt.Println("I'm a Golang Ninja")
		break
	}

	// SWITCH
	// Use switch statements to handle different cases.
	weapon = "Ninja Star"
	switch weapon {
	case "Ninja Star":
		fmt.Println("It's a Ninja Star!")
	case "Ninja Sword":
		fmt.Println("It's a Ninja Sword!")
	}
	powerLevel := 9001
	switch {
	case powerLevel > 9000:
		fmt.Println("It's over...NINE THOUSAND!!!")
	default:
		fmt.Println("It's a Baby Ninja")
	}

	// ARRAYS
	// Declare and initialize arrays using the var keyword.
	var evilNinjas [3]string
	fmt.Println(len(evilNinjas))
	evilNinjas[0] = "Johnny"
	fmt.Println(evilNinjas)
	fmt.Println(evilNinjas[0])
	fmt.Println(len(evilNinjas))
	moreEvilNinjas := [3]string{"Andy", "Tommy", "Bobby"}
	fmt.Println(moreEvilNinjas)
	var missionRewards [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			missionRewards[i][j] = i + j
		}
	}

	// SLICES
	// Declare and initialize slices using the var keyword.
	var evilNinjasSlice []string
	fmt.Println(len(evilNinjasSlice))
	evilNinjasSlice = append(evilNinjasSlice, "Tommy")
	fmt.Println(len(evilNinjasSlice))

	// MAPS
	// Declare and initialize maps using the make function.
	ninjaLevels := make(map[string]int)
	ninjaLevels["Johnny"] = 7
	ninjaLevels["Tommy"] = 13
	fmt.Println(ninjaLevels)
	fmt.Println(len(ninjaLevels))
	delete(ninjaLevels, "Johnny")
	fmt.Println(len(ninjaLevels))
	_, ok := ninjaLevels["Tommy"]
	fmt.Println(ok)
	moreNinjaLevels := map[string]int{"Bobby": 8, "Andy": 3}
	fmt.Println(moreNinjaLevels)

	// RANGE
	// Use range to iterate over slices and maps.
	evilNinjasRange := []string{"Tommy", "Johnny", "Andy"}
	for index, evilNinja := range evilNinjasRange {
		fmt.Println("Attacking target", index, evilNinja)
	}
	evilNinjasWithLevels := map[string]int{"Tommy": 2}
	for evilNinja, level := range evilNinjasWithLevels {
		fmt.Printf("%s -> %d\n", evilNinja, level)
	}

	// STRUCTS
	// Declare and initialize structs using the type keyword.
	type ninja struct {
		name  string
		level int
	}
	fmt.Println(ninja{name: "Bobby", level: 20})
	fmt.Println(ninja{name: "Andy", level: 30})
	fmt.Println(ninja{name: "Johnny"})
	tommy := ninja{name: "Tommy", level: 50}
	fmt.Println(tommy.level)
	tommy.level = 51

	// POINTERS
	// Declare and initialize pointers using the & operator.
	tommyPointer := &tommy
	johnnyPointer := &ninja{"Johnny", 0}
	var ninjaPointer *ninja = new(ninja)
	fmt.Println(tommyPointer, johnnyPointer, ninjaPointer)

	// INTERFACE
	// Declare and initialize interfaces using the type keyword.
	type Weapon interface {
		attack()
	}
	type ninjaStar struct{}
	type ninjaSword struct{}

	// Method for ninjaStar
	var _ = func(n ninjaStar) {
		fmt.Println("Throwing Ninja Star")
	}
}
	// Method for ninjaSword
	func (n ninjaSword) attack() {
		fmt.Println("Throwing Ninja Sword")
	}

	// Create a slice of Weapon interface
	weapons := []Weapon{
		ninjaStar{},
		ninjaSword{},
	}
	// Iterate over weapons and call attack method
	for _, weapon := range weapons {
		weapon.attack()
	}

	// FUNCTIONS
	// Declare and initialize functions using the func keyword.
	// Anonymous function to use a weapon
	useWeapon := func(ninja string, weapon string) string {
		return fmt.Sprintf(ninja + " is using " + weapon)
	}
	// Anonymous function to validate level
	isValidLevel := func(level int) (int, bool) {
		if level > 10 {
			return level, true
		}
		return level, false
	}
	// Variadic function to attack multiple targets
	attack := func(evilNinjas ...string) {
		for _, evilNinja := range evilNinjas {
			fmt.Println("Attacking target", evilNinja)
		}
	}
	usage := useWeapon("Tommy", "Ninja Star")
	level, valid := isValidLevel(11)
	fmt.Println(usage, level, valid)
	attack("Tommy", "Johnny")
	attack("Tommy", "Johnny", "Andy", "Bobby")
	evilNinjasAttack := []string{"Tommy", "Johnny", "Andy"}
	attack(evilNinjasAttack...)
	attackToo := attack
	attackToo(evilNinjasAttack...)
	// Immediately invoked function expression (IIFE)
	func() {
		fmt.Println("Attacking Evil Ninjas...")
	}()

	// GOROUTINES
	// Use goroutines to run functions concurrently.
	attackGoroutine := func(target string) {
		fmt.Println("Throwing ninja stars at", target)
	}
	go attackGoroutine("Tommy")
	time.Sleep(time.Second)

	// CHANNELS
	// Use channels to communicate between goroutines.
	attackChannel := func(target string, attacked chan bool) {
		time.Sleep(time.Second)
		fmt.Println("Throwing ninja stars at", target)
		attacked <- true
	}
	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attackChannel(evilNinja, smokeSignal)
	fmt.Println(<-smokeSignal)
	moreSmokeSignal := make(chan bool, 1)
	moreSmokeSignal <- true
	fmt.Println(<-moreSmokeSignal)
	close(moreSmokeSignal)
	for message := range moreSmokeSignal {
		fmt.Println(message)
	}
}