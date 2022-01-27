package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Value1:")
	inp1, _ := reader.ReadString('\n')
	float1, err1 := strconv.ParseFloat(strings.TrimSpace(inp1), 64)
	if err1 != nil {
		fmt.Println(err1)
		panic("Value1 must be a number")
	}
	fmt.Print("Enter the operation (add, sub, mul, div):")
	inp2, _ := reader.ReadString('\n')
	inp2 = strings.TrimSpace(inp2)
	fmt.Print("Enter Value2:")
	inp3, _ := reader.ReadString('\n')
	float2, err2 := strconv.ParseFloat(strings.TrimSpace(inp3), 64)
	if err2 != nil {
		fmt.Println(err2)
		panic("Value2 must be a number")
	}
	switch inp2 {
	case "add":
		fmt.Printf("Adding %f and %f\n", float1, float2)
		fmt.Printf("The Sum of %f and %f is %f", float1, float2, float1+float2)
	case "sub":
		fmt.Printf("Subtracting %f and %f\n", float1, float2)
		fmt.Printf("The Sum of %f and %f is %f", float1, float2, float1-float2)
	case "mul":
		fmt.Printf("Multiplying %v and %v\n", float1, float2)
		fmt.Printf("The product of %v and %v is %v", float1, float2, float1*float2)
	case "div":
		fmt.Printf("Dividing %v by %v\n", float1, float2)
		fmt.Printf("The division of %v and %v is %v", float1, float2, float1/float2)
	default:
		fmt.Println("Enter a valid operator")
	}

}
