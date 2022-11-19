package main

import "fmt"

func main() {
	var number1, number2 float64
	var operator string

	fmt.Print("Enter the first nuber : ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the Second nuber : ")
	fmt.Scanln(&number2)

	fmt.Print("Enter the Operator( + - * / ** ) : ")
	fmt.Scanln(&oprator)

	switch operator{
	case "+" :
		fmt.Printf("%.3f %s %.3f = %.3f," number1, operaator, number2, number1 + number2)

	case "-" :
		fmt.Printf("%.3f %s %.3f = %.3f," number1, operaator, number2, number1 - number2)

	case "*" :
		fmt.Printf("%.3f %s %.3f = %.3f," number1, operaator, number2, number1 * number2)

	case "**" :
		fmt.Printf("%.3f %s %.3f = %.3f," number1, operaator, number2, number1 ** number2)

	case "/" :
		if number2 == 0.0{
			fmt.Println("Divide by zero Situation")
		}else{
			fmt.Printf("%.3f %s %.3f = %.3f," number1, operaator, number2, number1 / number2)
		}
	default:
		fmt.Println("Invalid Operator")
	}
	
}
