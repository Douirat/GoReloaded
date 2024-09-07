package reload

import "fmt"

// Declare an object struct:
type String struct {
	str string
}

type Model interface {
	// ConvertBase(n, bf, bt string) string
	Filter()
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
	
	fmt.Println(slc_str)
}
