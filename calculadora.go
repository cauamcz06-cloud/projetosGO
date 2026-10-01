// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	println("coloque 2 numeros para somar(nesse formato 2+2)")
// 	var input string
// 	fmt.Scan(&input)

// 	// operation := strings.Split(input, "")

// 	result := getResult(operation)
// 	fmt.Printf("%s %s %s = %d", operation[0], operation[1], operation[2], result)

// 	println(operation[0])
// 	println(operation[1])
// 	println(operation[2])
// }
// func getResult(operation []string) int {
// 	num1, _ := strconv.Atoi(operation[0])
// 	num2, _ := strconv.Atoi(operation[2])

// 	switch operation[1] {
// 	case "+":
// 		return num1 + num2
// 	case "-":
// 		return num1 - num2
// 	case "*":
// 		return num1 * num2
// 	case "/":
// 		return num1 / num2
// 	default:
// 		panic("operador não valido")
// 	}

// }

// // var num1, num2 int
// // fmt.Scan(num1, num2)
// // println(num1, num2)
// // sum := num1 + num2
// // fmt.Printf("%d + %d = %d", num1, num2, sum)
