package Rover

import (
	"fmt"
	"strings"
)

type rover struct{
	xCoOrdinate int
	yCoOrdinate int
	direction int
}


func getDirectionInInteger(input string) int{
	switch input {
	case "N" :
		return 1
	case "E" :
		return 2
	case "S" :
		return 3
	case "W" :
		return 4
	}
	return 0
}

func getDirectionInString(input int) string{
	switch input {
	case 1 :
		return "N"
	case 2 :
		return "E"
	case 3 :
		return "S"
	case 4 :
		return "W"
	}
	return "N"
}

func printResults(rover *rover)  {
	fmt.Println(rover.xCoOrdinate, rover.yCoOrdinate, getDirectionInString(rover.direction))
}


func operateRover(rover *rover, input string) {
	inputs := strings.Split(input, "")
	for i := 0; i < len(inputs); i++ {
		if inputs[i] == "L" {
			turnLeft(rover)
			fmt.Println(rover)
		} else if inputs[i] == "R" {
			turnRight(rover)
			fmt.Println(rover)
		} else if inputs[i] == "M" {
			move(rover)
			fmt.Println(rover)
		} else {
			fmt.Println("Speak in Mars language, please!");
		}
	}
}

func move(rover *rover) {
	if rover.direction == 1 {
		rover.yCoOrdinate++
	} else if rover.direction == 2 {
		rover.xCoOrdinate++
	} else if rover.direction == 3 {
		rover.yCoOrdinate--
	} else if rover.direction == 4 {
		rover.xCoOrdinate--
	}
}

func turnRight(rover *rover) {
	if (rover.direction + 1) > 4 {
		rover.direction = 1
	} else{
		rover.direction = rover.direction + 1
	}
}

func turnLeft(rover *rover) {
	if (rover.direction - 1) < 1 {
		rover.direction = 4
	} else{
		rover.direction = rover.direction - 1
	}
}