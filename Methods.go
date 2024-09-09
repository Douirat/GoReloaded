package reload

import "fmt"


// A function to make sure if a character is a punctation mark:
func PunctuationMark(r rune) bool {
	return (r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';')
}

// Function to clean/clear my strings:
func CleanStrings(slc []string) []string {
	for i := 0; i < len(slc); i++ {

		if !PunctuationMark(rune(slc[i][0])) {
			continue
		}
		if len(slc[i]) > 1 {
			j := 0
			for _, v := range slc[i] {
				if PunctuationMark(v) && i-1 > 0 {
					slc[i-1] += string(v)
					j++

				} else if j < len(slc[i]) {
					slc[i] = slc[i][j:]
					j = 0
				}
			}
			if j >= len(slc[i]) {
				slc = append(slc[:i], slc[i+1:]...)
				continue
			}
		} else if len(slc[i]) == 1 {
			if PunctuationMark(rune(slc[i][0])) {
				if i > 1 {
					slc[i-1] += slc[i]
					slc = append(slc[:i], slc[i+1:]...)
				}
			}
		}
	}
	return slc
}

// Function to handle filter the single quotes:
func Quoting(slc []string) []string {
	indicator := 0
	sub_slc := []string{}
for i :=0; i<len(slc); i++ {
	if slc[i] == "'" {
		if indicator == 0 {
			fmt.Println("The indicqtor is 0")
			if i+1 <= len(slc)-1 {
				slc[i+1] = slc[i] + slc[i+1]
			}
			indicator = 1
		} else {
			fmt.Println("The indicqtor is 1")
			slc[i-1] += slc[i]
			indicator = 0
		}
	}
}
fmt.Println(slc)
for _, v := range slc {
	if v != "'" {
		sub_slc = append(sub_slc, v)
	}
}

return CleanStrings(sub_slc)
}