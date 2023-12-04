package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mainoper string

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var check string
	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(strings.TrimSpace(text), " ", "")
		check = checking(text)
		switch check {
		case "bad":
			fmt.Println("Некорректный ввод")
		case "arab":
			fmt.Println(arabculc(text))
		case "rim":
			fmt.Println(rimculc(text))
		}
	}
}

func checking(text string) string {
	var checksimb = map[string]string{
		"0": "arab", "1": "arab", "2": "arab", "3": "arab", "4": "arab", "5": "arab", "6": "arab",
		"7": "arab", "8": "arab", "9": "arab",
		"I": "rim", "V": "rim", "X": "rim",
		"+": "oper", "-": "oper", "*": "oper", "/": "oper",
	}
	var arab, rim, oper int
	if text == "" {
		return "bad"
	}
	toSimb := strings.Split(text, "")
	for _, v := range toSimb {
		var found = false
		for sv, st := range checksimb {
			if v == sv {
				found = true
				switch st {
				case "arab":
					arab++
				case "rim":
					rim++
				case "oper":
					oper++
					mainoper = sv
				}
			}
		}
		if found == false {
			return "bad"
		}
	}
	if (oper != 1) || (arab > 0 && rim > 0) {
		return "bad"
	}
	if arab > 0 {
		return "arab"
	} else {
		return "rim"
	}
}

func arabculc(text string) string {
	toParts := strings.Split(text, mainoper)
	if toParts[0] == "" || toParts[1] == "" {
		return "Ошибка: в выражении один операнд"
	}
	num1, _ := strconv.Atoi(toParts[0])
	num2, _ := strconv.Atoi(toParts[1])
	if num1 > 10 || num2 > 10 {
		return "Числа должны не должны быть больше 10"
	}
	switch mainoper {
	case "+":
		return strconv.Itoa(num1 + num2)
	case "-":
		return strconv.Itoa(num1 - num2)
	case "*":
		return strconv.Itoa(num1 * num2)
	case "/":
		if num2 == 0 {
			return "Ошибка: деление на 0"
		} else {
			return strconv.Itoa(num1 / num2)
		}
	}
	return ""
}

func rimculc(text string) string {
	var RimArab = map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
		"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
		"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
		"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
		"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
		"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
		"LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
		"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
		"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	}
	var num1, num2, result int
	toParts := strings.Split(text, mainoper)
	for rimNum, arabNum := range RimArab {
		if toParts[0] == rimNum {
			num1 = arabNum
		}
		if toParts[1] == rimNum {
			num2 = arabNum
		}
	}
	switch mainoper {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	if result < 1 {
		return "Римские цифры не могут быть меньше 1"
	} else {
		for rimNum, arabNum := range RimArab {
			if result == arabNum {
				return rimNum
			}
		}
	}
	return ""
}
