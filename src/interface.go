package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//type RoverOperations interface {
//	move()
//	turnLeft()
//	turnRight()
//	printResults()
//}
//
//type Rover struct{
//	xCoOrdinate int
//	yCoOrdinate int
//	direction int
//}
//
//func (rover Rover) move() {
//	if rover.direction == 1 {
//		rover.yCoOrdinate++
//	} else if rover.direction == 2 {
//		rover.xCoOrdinate++
//	} else if rover.direction == 3 {
//		rover.yCoOrdinate--
//	} else if rover.direction == 4 {
//		rover.xCoOrdinate--
//	}
//}
//
//func (rover Rover) printResults(){
//	fmt.Println(rover.xCoOrdinate, rover.yCoOrdinate, getDirectionInString(rover.direction))
//}
//
//func (rover Rover) turnRight() {
//	if (rover.direction + 1) > 4 {
//		rover.direction = 1
//	} else{
//		rover.direction = rover.direction + 1
//	}
//}
//
//func (rover Rover) turnLeft() {
//	if (rover.direction - 1) < 1 {
//		rover.direction = 4
//	} else{
//		rover.direction = rover.direction - 1
//	}
//}
//
//func getDirectionInInteger(input string) int{
//	switch input {
//	case "N" :
//		return 1
//	case "E" :
//		return 2
//	case "S" :
//		return 3
//	case "W" :
//		return 4
//	}
//	return 0
//}
//
//func getDirectionInString(input int) string{
//	switch input {
//	case 1 :
//		return "N"
//	case 2 :
//		return "E"
//	case 3 :
//		return "S"
//	case 4 :
//		return "W"
//	}
//	return "N"
//}
//
//func main(){
//	fmt.Println("Enter the X Y and Direction of First Rover Deployed Position")
//	var firstRoverXCoOrdinate int
//	var firstRoverYCoOrdinate int
//	var firstRoverDirection string
//	fmt.Scan(&firstRoverXCoOrdinate, &firstRoverYCoOrdinate, &firstRoverDirection)
//
//	fmt.Println("Enter the Series of instructions to move First rover in plateau")
//	var firstRoverInput string
//	fmt.Scanln(&firstRoverInput)
//
//	//fmt.Println("Enter the X Y and Direction of Second Rover Deployed Position")
//	//var secondRoverXCoOrdinate int
//	//var secondRoverYCoOrdinate int
//	//var secondRoverDirection string
//	//fmt.Scan(&secondRoverXCoOrdinate, &secondRoverYCoOrdinate, &secondRoverDirection)
//	//
//	//fmt.Println("Enter the Series of instructions to move rover in plateau")
//	//var secondRoverInput string
//	//fmt.Scan(&secondRoverInput)
//
//	var r1 RoverOperations
//
//	r1 = Rover{firstRoverXCoOrdinate, firstRoverYCoOrdinate, getDirectionInInteger(firstRoverDirection)}
//
//	inputs := strings.Split(firstRoverInput, "")
//	for i := 0; i < len(inputs); i++ {
//		if inputs[i] == "L" {
//			r1.turnLeft()
//		} else if inputs[i] == "R" {
//			r1.turnRight()
//		} else if inputs[i] == "M" {
//			r1.move()
//		} else {
//			fmt.Println("Speak in Mars language, please!");
//		}
//	}
//	r1.printResults()
//	//r2 := rover{secondRoverXCoOrdinate, secondRoverYCoOrdinate, getDirectionInInteger(secondRoverDirection)}
//	//operateRover(r2, secondRoverInput)
//	//printResults(r2)
//}