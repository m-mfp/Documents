package dices
 
import "math/rand"
import "fmt"

func RollDices(dices int)int{
	if dices == 0{
		dices = 1
	}
	
	result := 0
	for dices > 0 {
		// choose a random number from 1 to 6
		roll := rand.Intn(6) + 1
		fmt.Printf("You rolled %d!\n", roll)
		result += roll

		// add 1 to the dice variable
		dices -= 1
	}
	return result
}