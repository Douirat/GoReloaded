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
	HandleQuotes() []string
}

// Instantiate a new String object:
func NewString(s string) *String {
	return &String{str: s}
}

// // A function to filter my functions:
func (s *String) Filter() {
	required := s.HandleQuotes()
	fmt.Println(required)
}

func (s *String) HandleQuotes() []string {
	str := s.str
	sub_str := ""
	slc := []string{}
	// flag := 0

	for _, v := range str {
		if v == 39 {
			if sub_str != "" {
				slc = append(slc, sub_str)
				sub_str = ""
			}

			slc = append(slc, string(v))
			continue
		}
		if v == 32 {
			if sub_str != "" {
				slc = append(slc, sub_str)
				sub_str = ""
			}
			continue
		} else {
			sub_str += string(v)
		}
	}
	if sub_str != "" {
		slc = append(slc, sub_str)
		sub_str = ""
	}
	
	fmt.Println(Quoting(slc))
	return slc
}
