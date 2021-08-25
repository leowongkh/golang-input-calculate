package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter first operand: ")
	scanner.Scan()
	op1, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Print("Enter second operand: ")
	scanner.Scan()
	op2, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("%d + %d = %d\n", op1, op2, op1+op2)
	fmt.Printf("%d - %d = %d\n", op1, op2, op1-op2)
	fmt.Printf("%d * %d = %d\n", op1, op2, op1*op2)
	fmt.Printf("%d / %d = %d\n", op1, op2, op1/op2)
	fmt.Printf("%d %% %d = %d\n", op1, op2, op1%op2)
}
