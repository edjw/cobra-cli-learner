package numberfunctions

import (
	"fmt"
	"math"
)

// an incomplete package that turn a large number into a string with a fair bit of help from gpt4

type unit struct {
	name  string
	value float64
}

var units = []unit{
	{"quintillion", 1e18},
	{"quadrillion", 1e15},
	{"trillion", 1e12},
	{"billion", 1e9},
	{"million", 1e6},
	{"thousand", 1e3},
	{"hundred", 1e2},
}

var numberwords = map[int]string{
	0:          "zero",
	1:          "one",
	2:          "two",
	3:          "three",
	4:          "four",
	5:          "five",
	6:          "six",
	7:          "seven",
	8:          "eight",
	9:          "nine",
	10:         "ten",
	11:         "eleven",
	12:         "twelve",
	13:         "thirteen",
	14:         "fourteen",
	15:         "fifteen",
	16:         "sixteen",
	17:         "seventeen",
	18:         "eighteen",
	19:         "nineteen",
	20:         "twenty",
	30:         "thirty",
	40:         "forty",
	50:         "fifty",
	60:         "sixty",
	70:         "seventy",
	80:         "eighty",
	90:         "ninety",
	100:        "hundred",
	1000:       "thousand",
	1000000:    "million",
	1000000000: "billion",
}

func ConvertToWords(num int) string {
	switch {
	case num < 20:
		return numberwords[num]
	case num < 100:
		tens := num / 10 * 10
		remainder := num % 10
		result := numberwords[tens]
		if remainder > 0 {
			result += "-" + numberwords[remainder]
		}
		return result
	default:
		for _, unit := range units {
			if num >= int(unit.value) {
				multiplier := num / int(unit.value)
				remainder := num % int(unit.value)
				result := ConvertToWords(multiplier) + " " + unit.name
				if remainder > 0 {
					result += " " + ConvertToWords(remainder)
				}
				return result
			}
		}
		return "" // Ideally, some form of error case handling should go here.
	}
}

func FormatLargeNumber(num float64) string {
	for _, unit := range units {
		if num >= unit.value {
			number := num / unit.value
			intNum := int(math.Round(number))
			numberword := ConvertToWords(intNum)
			if unit.value > 1 {
				return numberword + " " + unit.name
			}
			return numberword
		}
	}

	return fmt.Sprintf("%.0f", num)
}
