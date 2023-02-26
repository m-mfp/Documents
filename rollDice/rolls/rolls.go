package main

import "fmt"
import "rollDice/dices"

func main(){
	var numDice int
	fmt.Println("How many dices would you like to roll?")
	fmt.Scanln(&numDice)
	rolls := dices.RollDices(numDice)
	fmt.Println(rolls)
}