package main
//
//import (
//	"fmt"
//	"strings"
//)
//
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
//func printResults(xCoOrdinate int, yCoOrdinate int, direction int)  {
//	fmt.Println(xCoOrdinate, yCoOrdinate, getDirectionInString(direction))
//}
//
//
//func operateRover(x int,y int, direction int, input string) {
//	inputs := strings.Split(input, "")
//	for i := 0; i < len(inputs); i++ {
//		if inputs[i] == "L" {
//			direction = turnLeft(direction)
//		} else if inputs[i] == "R" {
//			direction = turnRight(direction)
//		} else if inputs[i] == "M" {
//			x, y, direction = move(x, y, direction)
//		} else {
//			fmt.Println("Speak in Mars language, please!");
//		}
//	}
//	printResults(x, y, direction)
//}
//
//func move(xCoOrdinate int, yCoOrdinate int, direction int)(int, int, int) {
//	if direction == 1 {
//		yCoOrdinate++
//	} else if direction == 2 {
//		xCoOrdinate++
//	} else if direction == 3 {
//		yCoOrdinate--
//	} else if direction == 4 {
//		xCoOrdinate--
//	}
//	return xCoOrdinate, yCoOrdinate, direction
//}
//
//func turnRight(direction int) int{
//	if (direction + 1) > 4 {
//		return 1
//	} else{
//		return direction + 1
//	}
//}
//
//func turnLeft(direction int) int {
//	if (direction - 1) < 1 {
//		return 4
//	} else{
//		return direction - 1
//	}
//}
//
//
//
//func main(){
//	//fmt.Println("Enter the x y coordinates of the plateau")
//	//var xCoOrdinate int
//	//var yCoOrdinate int
//	//fmt.Scan(&xCoOrdinate, &yCoOrdinate)
//
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
//
//	operateRover(firstRoverXCoOrdinate, firstRoverYCoOrdinate, getDirectionInInteger(firstRoverDirection), firstRoverInput)
//	//r2 := rover{secondRoverXCoOrdinate, secondRoverYCoOrdinate, getDirectionInInteger(secondRoverDirection)}
//	//operateRover(r2, secondRoverInput)
//	//printResults(r2)
//
//}
//
//
