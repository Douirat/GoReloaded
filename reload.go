package reload

import (
	"fmt"
)

// Declare an interrface to handle the content of my files:
type File_Handler interface {
	Determine()
}

func display(arr []string) {
	for _, v := range arr {
		fmt.Println(v)
	}
	}


func (f *File) Determine() {
	instances := []string{}
	values := []string{}
	str := ""
	for _, v := range f.Content {
		if v != 40 && v != 41 {
			str += string(v)
		} else if v == 40 {
			if str != "" {
				values = append(values, str)
				str = ""
			}
	
			continue
		} else if v == 41 {
			if str != "" {
				instances = append(instances, str)
				str = ""
			}
			continue
		}
	}
	if str != "" {
		values = append(values, str)
		str = ""
	}

	
display(instances)
display(values)

for _, v := range values {
	var stringer Model = NewString(v)
	stringer.Filter()
}
}

