package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr := strings.Fields(Scan1()) // массив, который ввел пользователь
	if len(arr) != 3 {
		fmt.Println("Необходимо ввести два операнда и один оператор")
		return
	}
	var num1 = arr[0]
	var num2 = arr[2]
	if romanToInt(arr[0]) == 0 && romanToInt(arr[2]) != 0 || romanToInt(arr[0]) != 0 && romanToInt(arr[2]) == 0 {
		fmt.Println("Нельзя использовать одновременно 2 системы исчисления")
	}
	if romanToInt(arr[0]) == 0 && romanToInt(arr[2]) == 0 {
		var a, _ = strconv.Atoi(num1)
		var b, _ = strconv.Atoi(num2)
		if a >= 1 && a <= 10 && b >= 1 && b <= 10 {
			fmt.Println(calc(a, b, arr[1]))
		} else {
			fmt.Println("Несоответствие требованиям. Диапазон чисел от 1 до 10")
			return
		}
	}
	if romanToInt(arr[0]) != 0 && romanToInt(arr[2]) != 0 {
		var a = romanToInt(arr[0])
		var b = romanToInt(arr[2])
		if a >= 1 && a <= 10 && b >= 1 && b <= 10 {

			var result, _ = calc(a, b, arr[1])
			fmt.Println(Roman(result))

		} else {
			fmt.Println("Несоответствие требованиям. Диапазон чисел от I до X")
			return
		}
	}

}

func calc(num1, num2 int, act string) (int, error) {
	if act == "+" {
		return num1 + num2, nil
	} else if act == "-" {
		return num1 - num2, nil
	} else if act == "*" {
		return num1 * num2, nil
	} else if act == "/" {
		return num1 / num2, nil
	}
	return 0, errors.New("Невычесляемое значение")
}

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return strings.ToUpper(in.Text())
}

func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{

		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}
