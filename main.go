package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// DetectOperation find an operation (+ - * /)
func DetectOperation(expression []rune) (operation rune, err error) {
	count := 0
	for _, v := range expression {
		if string(v) == "+" || string(v) == "-" || string(v) == "*" || string(v) == "/" {
			count++
			operation = v
		}
	}
	if count == 1 {
		return operation, nil
	}
	if count == 0 {
		return 0, errors.New("Вывод ошибки, так как строка не является математической операцией")
	}
	if count > 1 {
		return 0, errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор")
	}
	return operation, nil
}

// FindNumberInString  Get number, numberOfType, err
func FindNumberInString(s string) (int, bool, error) {
	numArab, err := strconv.Atoi(s)
	if err == nil {
		return numArab, true, nil
	}

	numRoman, err := RomanToArab(s)
	if err == nil {
		return numRoman, false, nil
	}

	return 0, false, err
}

// RomanToArab Convert roman to arabic
func RomanToArab(s string) (result int, err error) {
	var RomeAlphabet = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	if strings.Contains(strings.Join(RomeAlphabet, ","), s) == false {
		fmt.Println("\"not a number\"")
		return
	}

	Roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
	}
	lastElement := s[len(s)-1 : len(s)]
	result = Roman[lastElement]
	for i := len(s) - 1; i > 0; i-- {
		if Roman[s[i:i+1]] <= Roman[s[i-1:i]] {
			result += Roman[s[i-1:i]]
		} else {
			result -= Roman[s[i-1:i]]
		}
	}
	return result, nil
}

// Calc get result
func Calc(left, right int, sign int32) (result int, err error) {
	if left > 10 || right > 10 {
		return 0, errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор")
	}
	switch string(sign) {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right != 0 {
			return left / right, nil
		}
	default:
		return 0, errors.New("Вывод ошибки, так как строка не является математической операцией")

	}

	return 0, err
}

// ArabToRoman Convert arabic to roman
func ArabToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}
	return roman
}

func main() {
	fmt.Print("Введите выражение: ")
	in := bufio.NewReader(os.Stdin)
	var s string
	_, err := fmt.Fscan(in, &s)
	if err != nil {
		log.Fatal(err)
	}
	str := []rune(s) //Convert String to []Rune

	//Detect operation
	sign, err := DetectOperation(str) //sign это знак + - * /
	if err != nil {
		fmt.Println(err)
		return
	}

	convert := strings.Split(string(str), string(sign)) // Split
	//Find the left number
	leftNumber, leftNumType, err := FindNumberInString(strings.ToUpper(convert[0]))
	//Find the right number
	rightNumber, rightNumType, err := FindNumberInString(strings.ToUpper(convert[1]))

	//Checking the type of left and right digits
	if leftNumType != rightNumType {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	}
	//Get result Calc
	result, err := Calc(leftNumber, rightNumber, sign)
	if err != nil {
		return
	}

	if result <= 0 {
		fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		return
	}

	if leftNumType == false {
		fmt.Println(ArabToRoman(result))
		return
	}
	fmt.Println(result)
}
