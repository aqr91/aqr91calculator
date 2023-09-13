package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Calculator can perform +, -, * and / operations with two numbers : a + b, a - b, a * b, a / b\nNumbers can be either Arabic or Roman, but not a both combination")
	expression := bufio.NewReader(os.Stdin) // buff for string input
	fmt.Println("Input expression:")
	str, _ := expression.ReadString('\n') //reading till user input enter
	parts := strings.Split(str, " ")      //Split string to parts by space

	if len(parts) != 3 { //Checking if the input is correct
		panic("Invalid input format")
		return
		os.Exit(1)
	}
	fmt.Println("Input string:", str)

	firstnumber := parts[0]
	oper := parts[1]
	oper = strings.TrimSpace(oper) //space trimming for correct working
	secondnumber := parts[2]
	secondnumber = strings.TrimSpace(secondnumber)
	romNumbersStrToTen := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arabNumbersStr := []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var num1, romnum1, num2, romnum2, result, i int
	for i = 0; i < 11; i++ { //Checking, is first number meets req`s
		if firstnumber == romNumbersStrToTen[i] {
			num1 = i
			romnum1 = i
			break
		} else if firstnumber == arabNumbersStr[i] {
			num1 = i
			romnum1 = -1 * i
			break
		}
	}
	if i == 11 {
		panic("First variable does not meet requirements ")
		return
		os.Exit(2)
	}
	for i = 0; i < 11; i++ { //Same for second
		if secondnumber == romNumbersStrToTen[i] {
			num2 = i
			romnum2 = i
			break
		} else if secondnumber == arabNumbersStr[i] {
			num2 = i
			romnum2 = -1 * i
			break
		}
	}
	if i == 11 {
		panic("Second variable does not meet requirements ")
		return
		os.Exit(3)
	}
	if oper == "+" { //Checking operator for req`s
	} else if oper == "-" {
	} else if oper == "*" {
	} else if oper == "/" {
	} else {
		panic("Operator does not meet requirements ")
		return
		os.Exit(4)
	}
	if romnum1 < 0 && romnum2 > 0 || romnum1 > 0 && romnum2 < 0 { //Checking number systems
		panic("Different number systems used")
		return
		os.Exit(5)
	}

	if oper == "+" { //calculation
		result = num1 + num2
	} else if oper == "-" {
		result = num1 - num2 //Checking is result for rom system is < 1 for panic call
		if romnum1 > 0 && romnum2 > 0 && result < 1 {
			panic("Rom number cant be < 1")
			return
			os.Exit(6)
		} //Rom number cant be < 1
		result = num1 - num2
	} else if oper == "*" {
		result = num1 * num2
	} else if oper == "/" {
		result = num1 / num2
		if romnum1 > 0 && romnum2 > 0 && result < 1 { //same check
			panic("Rom number cant be < 1")
			return
			os.Exit(6)
		}
	}
	romNumbersStr := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XVLI", "XVLII", "XVLIII", "XVLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LVXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

	if romnum1 < 0 && romnum2 < 0 { //if user input rom nums, convert result to rom numbers and output
		fmt.Println("\n Result:", result)
	} else {
		fmt.Println("\n Result:", romNumbersStr[result]) //if user input arab nums - output result

	}
}
