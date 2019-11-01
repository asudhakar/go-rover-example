package main

import (
	"fmt"
)


func main(){
	fmt.Println("Enter the x y coordinates of the plateau")
	var xCoOrdinate int
	var yCoOrdinate int
	fmt.Scan(&xCoOrdinate, &yCoOrdinate)

	fmt.Println("Enter the X Y and Direction of First Rover Deployed Position")
	var firstRoverXCoOrdinate int
	var firstRoverYCoOrdinate int
	var firstRoverDirection string
	fmt.Scan(&firstRoverXCoOrdinate, &firstRoverYCoOrdinate, &firstRoverDirection)

	fmt.Println("Enter the Series of instructions to move First rover in plateau")
	var firstRoverInput string
	fmt.Scanln(&firstRoverInput)

	fmt.Println("Enter the X Y and Direction of Second Rover Deployed Position")
	var secondRoverXCoOrdinate int
	var secondRoverYCoOrdinate int
	var secondRoverDirection string
	fmt.Scan(&secondRoverXCoOrdinate, &secondRoverYCoOrdinate, &secondRoverDirection)

	fmt.Println("Enter the Series of instructions to move rover in plateau")
	var secondRoverInput string
	fmt.Scan(&secondRoverInput)

	r1 := rover{firstRoverXCoOrdinate, firstRoverYCoOrdinate, getDirectionInInteger(firstRoverDirection)}

	operateRover(&r1, firstRoverInput)
	printResults(&r1)
	r2 := rover{secondRoverXCoOrdinate, secondRoverYCoOrdinate, getDirectionInInteger(secondRoverDirection)}
	operateRover(&r2, secondRoverInput)
	printResults(&r2)

}

