package main

import (
	"fmt"
)

// available ingredients
var waterAmount int = 400
var milkAmount int = 540
var coffeeAmount int = 120
var disposableCups int = 9
var moneyAmount int = 550
var coffeeCups int
var action string
var option int

var waterPtr = &waterAmount
var milkPtr = &milkAmount
var coffeePtr = &coffeeAmount
var disposableCupsPtr = &disposableCups
var moneyPtr = &moneyAmount

var addedWater int
var addedMilk int
var addedCoffee int
var addedCups int

// get ingredient functions
func getWater() int {
	return *waterPtr
}

func getMilk() int {
	return *milkPtr
}

func getCoffee() int {
	return *coffeePtr
}

func getDisposableCups() int {
	return *disposableCupsPtr
}

func getMoney() int {
	return *moneyPtr
}

func getState() {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d of water\n", getWater())
	fmt.Printf("%d of milk\n", getMilk())
	fmt.Printf("%d of coffee beans\n", getCoffee())
	fmt.Printf("%d of dusposable cups\n", getDisposableCups())
	fmt.Printf("%d of money\n", getMoney())
}

// calculating ingredient functions
func calculateAmountOfWater(cups int) int {
	return cups * 200
}

func calculateAmountOfMilk(cups int) int {
	return cups * 50
}

func calculateAmountOfCoffee(cups int) int {
	return cups * 15
}

// get min value
func getMin(v1,v2,v3 int) int {
	if v1 <= v2 {
		if v1 <= v3 {
			return v1
		} else {
			return v3
		}
	} else if v2 <= v3 {
		return v2
	} else {
		return v3
	}
}

// check availability
func checkAvailability(option int) bool{
	switch option {
	case 1:
		return *waterPtr >= 250 && *coffeePtr >= 16 && *disposableCupsPtr >= 1
	case 2:
		return *waterPtr >= 350 && *coffeePtr >= 20 && *milkPtr >= 75 && *disposableCupsPtr >= 1
	case 3:
		return *waterPtr >= 200 && *coffeePtr >= 12 && *milkPtr >= 100 && *disposableCupsPtr >= 1
	default:
		return false
	}
}

// handle buy
func handleBuy(option int) {
	switch option {
	case 1:
		availability := checkAvailability(1)
		if availability {
			*waterPtr -= 250
			*coffeePtr -= 16
			*disposableCupsPtr -= 1
			*moneyPtr += 4
		}
	case 2:
		availability := checkAvailability(2)
		if availability {
			*waterPtr -= 350
			*milkPtr -= 75
			*coffeePtr -= 20
			*disposableCupsPtr -= 1
			*moneyPtr += 7
		}
	case 3:
		availability := checkAvailability(3)
		if availability {
			*waterPtr -= 200
			*milkPtr -= 100
			*coffeePtr -= 12
			*disposableCupsPtr -= 1
			*moneyPtr += 6
		}
	default:
		fmt.Println("You pressed wrong option!")
	}
}

// handle fill
func handleFill() {
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&addedWater)
	*waterPtr += addedWater
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&addedMilk)
	*milkPtr += addedMilk
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&addedCoffee)
	*coffeePtr += addedCoffee
	fmt.Println("Write how many disposable coffee cups you want to add:")
	fmt.Scan(&addedCups)
	*disposableCupsPtr += addedCups
}

// handle take
func handleTake() {
	moneyToTake := getMoney()
	*moneyPtr -= moneyToTake
	fmt.Printf("I gave you $%d\n", moneyToTake)
}

// handle user action
func handleAction() {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)
	fmt.Println("")

	switch action {
	case "remaining":
		getState()
	case "exit":
		break
	case "buy":
		fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
		fmt.Scan(&option)
		handleBuy(option)
	case "fill":
		handleFill()
	case "take":
		handleTake()
	default:
		fmt.Println("The option is not available")
	}
}

// calculate extra cups
func calculateExtraCups(water, milk, coffee int) int {
	cupOfWater := int(water / 200)
	cupOfMilk := int(milk / 50)
	cupOfCoffee := int(coffee / 15)

	extraCups := getMin(cupOfWater, cupOfMilk, cupOfCoffee)
	return extraCups
}

// calculate available coffee cups
func calculateCoffeeCups(cups int) {
	requiredWater := calculateAmountOfWater(cups)
	requiredMilk := calculateAmountOfMilk(cups)
	requiredCoffee := calculateAmountOfCoffee(cups)

	switch {
		case requiredWater > waterAmount || requiredMilk > milkAmount || requiredCoffee > coffeeAmount:
			availableCups := calculateExtraCups(waterAmount, milkAmount, coffeeAmount)
			fmt.Printf("No, I can make only %d cups of coffee\n", availableCups)
		case requiredWater < waterAmount || requiredMilk < milkAmount || requiredCoffee < coffeeAmount:
			leftOverWater := waterAmount - requiredWater
			leftOverMilk := milkAmount - requiredMilk
			leftOverCoffee := coffeeAmount - requiredCoffee

			extraCups := calculateExtraCups(leftOverWater, leftOverMilk, leftOverCoffee)
			if extraCups == 0 {
				fmt.Println("Yes, I can make that amount of coffee")
			} else {
				fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", extraCups)
			}
		default:
			fmt.Println("Yes, I can make that amount of coffee")
	}
}

func main() {
	handleAction()
}
