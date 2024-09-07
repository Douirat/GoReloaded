package reload

import (
	"fmt"
)

// Declare an object struct:
type String struct {
	str string
}

type Model interface {
	// ConvertBase(n, bf, bt string) string
	Filter()
}

// Function to clean/clear my strings:
func cleanStrs(slc []string) []string {
	// str := ""
	for i := 0; i < len(slc); i++ {
		if len(slc[i]) == 1 {
			if PunctuationMark(rune(slc[i][0])) {
				slc[i-1] += slc[i]
				slc = append(slc[:i], slc[i+1:]...)

			}
		} else if len(slc[i]) > 1 {
			j := 0
			for _, v := range slc[i] {
				if PunctuationMark(v) {
					slc[i-1] += string(v)
					j++
					
				} else if j < len(slc[i]) {
					slc[i] = slc[i][j:]
					j=0
				}
			}
			if j >= len(slc[i]) {
				slc = append(slc[:i], slc[i+1:]...)
				continue
			}
		}
	}
	return slc
}

// A function to make sure if a character is a punctation mark:
func PunctuationMark(r rune) bool {
	return (r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';')
}

// Instantiate a new String object:
func NewString(s string) *String {
	return &String{str: s}
}

// Filter the content of my string:
func (s *String) Filter() {
	str := ""
	slc_str := []string{}
	for _, v := range s.str {
		if v != ' ' {
			str += string(v)
		} else {
			if str != "" {
				slc_str = append(slc_str, str)
				str = ""
			}
		}
	}
	if str != "" {
		slc_str = append(slc_str, str)
		str = ""
	}
	slc := cleanStrs(slc_str)
	fmt.Println(slc)
}
