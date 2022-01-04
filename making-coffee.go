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

// retrieving ingredient functions
func getWater() int {
	return waterAmount
}

func getMilk() int {
	return milkAmount
}

func getCoffee() int {
	return coffeeAmount
}


func main() {
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&waterAmount)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milkAmount)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&coffeeAmount)
	fmt.Println("Write how many cups of coffee you will need:")

	fmt.Scan(&coffeeCups)

	calculateCoffeeCups(coffeeCups)
}
