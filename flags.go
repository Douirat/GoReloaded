package Reload

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

// Convert from a base to another:
func ConvertBase(nbr, from, to string) string {
	n := 0
	for _, v := range nbr {
		for i := range from {
			if rune(from[i]) == v {
				n *= len(from)
				n += i
			}
		}
	}
	str := ""
	for n > 0 {
		str = string(rune(to[n%len(to)])) + str
		n /= len(to)
	}
	return str
}

// make sure the string is a flag:
// A function to capitalize a string:
func Capitalize(word string) string {
	var str string
	changed := false
	if isLowerCase(rune(word[0])) && isLetter(rune(word[0])) {
		str += string(rune(word[0]) - 32)
		changed = true
	}
	if changed {
		str += word[1:]
		return str
	}
	return word
}

// A function to convert a string to upper case:
func ToUpperCase(str string) string {
	sub_str := ""
	for _, v := range str {
		if isLowerCase(v) {
			sub_str += string(v - 32)
			continue
		}
		sub_str += string(v)
	}
	return sub_str
}

// A function to convert a string to lowercase:
func ToLowerCase(str string) string {
	sub_str := ""
	for _, v := range str {
		if isUpperCase(v) {
			sub_str += string(v + 32)
			continue
		}
		sub_str += string(v)
	}
	return sub_str
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
	for i, word := range slc {
		switch word {
		case "(cap,":
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
			}

		case "(low,":
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
			}

		case "(up,":
			if i < len(slc)-1 && rune(slc[i+1][len(slc[i+1])-1]) == 41 && isNumber(rune(slc[i+1][0])) {
				ind := Atoi(slc[i+1][:len(slc[i+1])-1])
				// fmt.Println(ind)
				j := i
				for j > 0 && ind > 0 {
					slc[j-1] = ToUpperCase(slc[j-1])
					ind--
					j--
				}
				slc = append(slc[:i], slc[i+1:]...)
				slc = append(slc[:i], slc[i+1:]...)
			}
		case "(low)":
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = ToLowerCase(slc[i-1])
				}
				slc = append(slc[:i], slc[i+1:]...)
			}

		case "(up)":
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = ToUpperCase(slc[i-1])
				}
				slc = append(slc[:i], slc[i+1:]...)
			}

		case "(cap)":
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = Capitalize(slc[i-1])
				}
				slc = append(slc[:i], slc[i+1:]...)
			}
		case "(hex)":
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = ConvertBase(slc[i-1], "0123456789ABCDEF", "0123456789")
				}
				slc = append(slc[:i], slc[i+1:]...)
			}
		case "(bin)":
			if i < len(slc) {
				if i > 0 {
					slc[i-1] = ConvertBase(slc[i-1], "01", "0123456789")
				}
				slc = append(slc[:i], slc[i+1:]...)
			}

		}
	}
	// fmt.Println(slc)
	result := new(string)
	*result = slc[0]
	for i, sub := range slc {
		if i != 0 {
			*result += " " + sub
		}
	}
	line.String = *result
}
