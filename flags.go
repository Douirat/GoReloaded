package Reload

import (
	"fmt"
	"unicode"
)

// => functions to determine the type of the charcters:
// Is lower case?
func isLowerCase(r rune) bool {
	return r >= 97 && r <= 122
}

// Is upper case:
func isUpperCase(r rune) bool {
	return r >= 65 && r <= 90
}

// Is letter:
func isLetter(r rune) bool {
	return r >= 65 && r <= 90 || r >= 97 && r <= 122
}

// Is alfa:
// func isAlfa(r rune) bool {
// 	return r >= 65 && r <= 90 || r >= 97 && r <= 122 || r >= 48 && r <= 57
// }

// Is number:
func isNumber(r rune) bool {
	return r >= 48 && r <= 57
}

// A function to convert an ascii / string to integer values
func Atoi(str string) int {
	x, y := 0, 1
	for i, v := range str {
		if isNumber(v) {
			x = x*10 + (int(v - 48))
		} else if i == 0 && v == 45 {
			x *= -1
		}
	}
	return x * y
}

// Check the base if valid:
func IsValidBase(n, base string) bool {
	Chosen := map[rune]bool{}
	if base == "01" {
		Chosen = map[rune]bool{48: true, 49: true}
	} else {
		Chosen = map[rune]bool{48: true, 49: true, 50: true, 51: true, 52: true, 53: true, 54: true, 55: true, 56: true, 57: true, 65: true, 66: true, 67: true, 68: true, 69: true, 70: true}
	}
	for _, b := range n {
		if !Chosen[b] {
			return false
		}
	}
	return true
}

// Check the form of my string if upper case or not:
func UpperCase(word string) bool {
	// Convert the word to runes to handle Unicode characters correctly
	runes := []rune(word)
	for _, r := range runes {
		// If any rune is not uppercase and not a space, return false
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Convert from a base to another:
func ConvertBase(nbr, from, to string) string {
	var num string
	negative := false

	if rune(nbr[0]) == 45 {
		negative = true
		num = nbr[1:]
	} else {
		num = nbr
	}

	if rune(nbr[0]) == 43 {
		num = nbr[1:]
	}

	if !IsValidBase(ToUpperCase(num), from) {
		return nbr
	}
	num = ToUpperCase(num)
	fmt.Println(num)
	n := 0
	for _, v := range num {
		for i := range from {
			if rune(from[i]) == v {
				n *= len(from)
				n += i
			}
		}
	}
	if n == 0 {
		return "0"
	}
	str := ""
	for n > 0 {
		str = string(rune(to[n%len(to)])) + str
		n /= len(to)
	}
	if negative {
		str = "-" + str
	}
	return str
}

// make sure the string is a flag:
// A function to capitalize a string:
// func Capitalize(word string) string {
// 	str := ""
// 	changed := false
// 	edge := 1
// for i:=0; i<len(word); i++ {
// 	if isLowerCase(rune(word[i])) && isLetter(rune(word[i])) {
// 		str += string(rune(word[i]) - 32)
// 		changed = true
// 	}
// 	if changed {
// 		break
// 	}
// 	if isLetter(rune(word[i])) {
// 		return word
// 	}
// 	str += string(rune(word[i]))
// 	edge++
// }
// if changed {
// 	if edge < len(word) {
// 	 str += word[edge:]
// 	}
// 	return str
// }
// return word
// }

func Capitalize(word string) string {
	// Convert the word to runes to handle Unicode characters correctly
	runes := []rune(word)
	if len(runes) == 0 {
		return word
	}
	// Capitalize the first rune
	runes[0] = unicode.ToUpper(runes[0])
	// Lowercase the rest of the runes
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// A function to convert a string to upper case:
func ToUpperCase(word string) string {
	// Convert the word to runes to handle Unicode characters correctly
	runes := []rune(word)
	// Convert all runes to lowercase
	for i, r := range runes {
		runes[i] = unicode.ToUpper(r)
	}
	return string(runes)
}

// A function to convert a string to lowercase:
func ToLowerCase(word string) string {
	// Convert the word to runes to handle Unicode characters correctly
	runes := []rune(word)
	// Convert all runes to lowercase
	for i, r := range runes {
		runes[i] = unicode.ToLower(r)
	}
	return string(runes)
}

// make changes on a string based on the adequat flag:
func (line *Line) HandleFlags() {
	slc := []string{}
	s := ""
	for _, char := range line.String {
		if char == 32 {
			if s != "" {
				slc = append(slc, s)
				s = ""
			}
			continue
		}
		s += string(char)
	}
	if s != "" {
		slc = append(slc, s)
		s = ""
	}
	for i := 0; i < len(slc); i++ {
		switch slc[i] {
		case "(cap,":
			if len(slc) == 1 {
				break
			}
			if i < len(slc)-1 && rune(slc[i+1][len(slc[i+1])-1]) == 41 && isNumber(rune(slc[i+1][0])) {
				ind := Atoi(slc[i+1][:len(slc[i+1])-1])
				j := i
				for j > 0 && ind > 0 {
					slc[j-1] = Capitalize(slc[j-1])
					ind--
					j--
				}
				slc = append(slc[:i], slc[i+1:]...)
				slc = append(slc[:i], slc[i+1:]...)
				i = 0

			}

		case "(low,":
			if len(slc) == 1 {
				break
			}
			if i < len(slc)-1 && rune(slc[i+1][len(slc[i+1])-1]) == 41 && isNumber(rune(slc[i+1][0])) {
				ind := Atoi(slc[i+1][:len(slc[i+1])-1])
				// fmt.Println(ind)
				j := i
				for j > 0 && ind > 0 {
					slc[j-1] = ToLowerCase(slc[j-1])
					ind--
					j--
				}
				slc = append(slc[:i], slc[i+1:]...)
				slc = append(slc[:i], slc[i+1:]...)
				i = 0

			}

		case "(up,":
			if len(slc) == 1 {
				break
			}
			if i < len(slc)-1 && rune(slc[i+1][len(slc[i+1])-1]) == 41 && isNumber(rune(slc[i+1][0])) {
				ind := Atoi(slc[i+1][:len(slc[i+1])-1])
				j := i
				for j > 0 && ind > 0 {
					slc[j-1] = ToUpperCase(slc[j-1])
					ind--
					j--
				}
				slc = append(slc[:i], slc[i+1:]...)
				slc = append(slc[:i], slc[i+1:]...)
				i = 0

			}
		case "(low)":
			if len(slc) == 1 {
				slc = []string{}
				break
			}
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = ToLowerCase(slc[i-1])
				}
				slc = append(slc[:i], slc[i+1:]...)
				i = 0

			}

		case "(up)":
			if len(slc) == 1 {
				slc = []string{}
				break
			}
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = ToUpperCase(slc[i-1])
				}
				slc = append(slc[:i], slc[i+1:]...)
				i = 0

			}

		case "(cap)":
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = Capitalize(slc[i-1])
				}
				slc = append(slc[:i], slc[i+1:]...)
				i = 0

			}
		case "(hex)":
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = ConvertBase(slc[i-1], "0123456789ABCDEF", "0123456789")
				}
				slc = append(slc[:i], slc[i+1:]...)
				i = 0
			}
		case "(bin)":
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = ConvertBase(slc[i-1], "01", "0123456789")
				}
				slc = append(slc[:i], slc[i+1:]...)
				i = 0
			}

		}
	}
	if len(slc) == 0 {
		line.String = ""
		return
	}
	result := new(string)
	*result = slc[0]
	for i, sub := range slc {
		if i != 0 {
			*result += " " + sub
		}
	}
	line.String = *result
}
