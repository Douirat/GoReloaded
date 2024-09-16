package Reload

import "fmt"

// Declare a string object:
type Line struct {
	String string // this field represents a line of the text
	Next   *Line  // This is a pointer to the next line un a queue of lines
}

// Declare a text object to represent the whole text following the FIFO principle:
type Text struct {
	First *Line // A pointer to the first line in the text.
	Last  *Line // Apointer to the last line in the text.
	Size  int   // To keep track of how many lines i have in my text
}

// Daclare an input object:
type Input struct {
	String string
}

// Instantiate a new line:
func NewLine(str string) *Line {
	new_line := new(Line)
	new_line.String = str
	new_line.Next = nil
	return new_line
}

// Instantiate a new Text object:
func NewText() *Text {
	new_text := new(Text)
	new_text.First = nil
	new_text.Last = nil
	new_text.Size = 0
	return new_text
}

// Instantiate a new input object:
func NewInput(text string) *Input {
	input := new(Input)
	input.String = text
	return input
}

// Is my queue empty:
func (text *Text) IsEmpty() bool {
	return text.First == nil
}

// Add a line to the Queue based on its location in the list:
func (text *Text) Enqueue(str string) {
	new_line := NewLine(str)
	if text.Last == nil {
		// if the queue is empty:
		text.First = new_line
		text.Last = new_line
	} else {
		// append to the queue:
		text.Last.Next = new_line
		text.Last = new_line
	}
	text.Size++
	str = ""
}

// Deque my text into a single string:
func (line *Line) Dequeue() string {
	if line.Next == nil {
		return line.String
	}
	 return line.String + "\n" +line.Next.Dequeue()
}

// Divide the input string into lines to make it ease the manipulation of the correction:
func (text *Text) DivideInput(input *Input) {
	str := ""
	for _, char := range input.String {
		if char == 10 {
			if str != "" {
				text.Enqueue(str)
				str = ""
			}
			continue
		}
		str += string(char)
	}
	if str != "" {
		text.Enqueue(str)
	}
}

// Trim excesses in white spaces:
func (line *Line) TrimExcessSpaces() {
	str := ""
	slc := []string{}
	for _, char := range line.String {
		if char == 32 || char == 9 {
			if str != "" {
				slc = append(slc, str)
				str = ""
			}
			continue
		}
		str += string(char)
	}
	if str != "" {
		slc = append(slc, str)
		str = ""
	}
	
	if len(slc) == 0 {
		line.String = ""
		return
	}
	
	result := slc[0] 
	for i, word := range slc {
		if i != 0 {
			result += " " + word
		}
	}
	fmt.Println(result)
	line.String = result
}

// A function to make sure the character is a punctuation mark:
func isPunctuationMark(r rune) bool {
	marks := ".,!?:;"
	for _, c := range marks {
		if r == c {
			return true
		}
	}
	return false
}

func (line *Line) handlePunctution() {
	s := ""
	slc := []string{}
	for _, char := range line.String {
		if char == 32 {
			if s != "" {
				slc = append(slc, s)
				s = ""
			}
			continue
		}
		if isPunctuationMark(char) {
			if s == "" {
				if len(slc) > 0 {
					slc[len(slc)-1] += string(char)
					continue
				}
			} else if s != "" {
				s += string(char)
				continue
			}
		}
		s += string(char)
	}
	if s != "" {
		slc = append(slc, s)
		s = ""
	}
	if len(slc) == 0 {
		line.String = ""
		return
	}
	result := new(string)
	*result = slc[0]
	for i, word := range slc {
		if i != 0 {
			*result += " " + word
		}
	}

	line.String = *result
}

// A function to make sure if a character is a punctation mark:
// Relates to the organize quotes function
func PunctuationMark(r rune) bool {
	return (r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';')
}

// A func tion to make sure either a char is an alfanumeric or not:
// Relates to the organize quotes function
func IsAlfa(r rune) bool {
	return r >= 65 && r <= 90 || r >= 97 && r <= 122 || r >= 48 && r <= 57 || PunctuationMark(r)
}

// A function to handle the quotes
func (line *Line) organizeQuotes() {
	str := line.String
	s := ""
	within := false
	pass := false
	for i, char := range str {
		if char == 39 {
			if !within {
				within = true
				if i > 0 && IsAlfa(rune(str[i-1])) && i < len(str)-1 && IsAlfa(rune(str[i+1])) {
					within = false
					s += string(char)
					continue
				}
				// fmt.Println("we are within the quotes")
			} else {
				within = false
				if i > 0 && IsAlfa(rune(str[i-1])) && i < len(str)-1 && IsAlfa(rune(str[i+1])) {
					within = true
					s += string(char)
					continue
				}
				// fmt.Println("we are getting out!!!")
			}
		}
		if within {
			if char == 32 && IsAlfa(rune(str[i-1])) && IsAlfa(rune(str[i+1])) {
				pass = true
			} else if char == 32 && !IsAlfa(rune(str[i-1])) && IsAlfa(rune(str[i+1])) {
				pass = false
			} else if char == 32 && IsAlfa(rune(str[i-1])) && !IsAlfa(rune(str[i+1])) {
				pass = false
			}
		}
		// if within && pass && char == 32 {
		// 	// fmt.Println(char)
		// } else if within && char == 32 {
		// 	// fmt.Println("-----")
		// } else if within {
		// 	// fmt.Println("<><><><>")
		// }
		if within && !pass && char == 32 {
			continue
		}
		s += string(char)
	}
	line.String = s
}


// // Trim excesses in white spaces:
// func  trimStr(s string) string {
// 	// if s == " " || s == ""{
// 	// 	return ""
// 	// }
// 	if s == " " {
// 		return " "
// 	} else if s == "" {
// 		return ""
// 	}
// 	str := ""
// 	slc := []string{}
// 	for _, char := range s {
// 		if char == 32 || char == 9 {
// 			if str != "" {
// 				slc = append(slc, str)
// 				str = ""
// 			}
// 			continue
// 		}
// 		str += string(char)
// 	}
// 	if str != "" {
// 		slc = append(slc, str)
// 		str = ""
// 	}
// 	result := slc[0]
// 	for i, word := range slc {
// 		if i != 0 {
// 			result += " " + word
// 		}
// 	}
// 	return result
// }


// func (line *Line) organizeQuotes() {
// s := trimStr(line.String)
// fmt.Println(s)
// slc := []string{}
// Quoted := false
// start := 0
// end := 0
// for i, char := range s {
// 	if char == 39 {
// 		if !Quoted {
// 			Quoted = true
// 			// if start == 0 && i == 0 {
// 			// 	start = i
// 			// 	continue belief, it was         the epoch of incredulity,             it was the season of Light, it was the season
 
 
 
// 			// }
// 			start = i
// 			slc = append(slc, s[end:start])
// 			continue
// 		} else if Quoted {
// 			Quoted = false
// 			end = i+1
// 			slc = append(slc, s[start:end])
// 		}
// 	}
// }
// if end < len(s) {
// 	slc = append(slc, s[end:])
// }


// for i, word := range slc{
// if len(word) > 1 {
// 	if rune(word[0]) == 39 && rune(word[len(word)-1]) == 39 {
// 		slc[i] = "'" + trimStr(word[1:len(word)-1]) + "'"
// 	}
// }
// }
// result := slc[0]
// for j, v := range  slc {
// 	if j != 0 {
// 		result += v
// 	}
// }
// line.String = result
// }

// A function to make sure my method is a vowel:
// Related to a to an method
func IsVowel(r rune) bool {
	runes := "aAeEiIoOuU"
	for _, v := range runes {
		if r == v {
			return true
		}
	}
	return false
}

// A function to make sure the if it is a special case (H):
// Related to a to an method:
func IsSpecial(org_str string) bool {
	specials := []string{"hour", "honor", "honest", "heir"}
	for _, str := range specials {
		if org_str == str {
			return true
		}
	}
	return false
}

// Convert a/A to an/An if the next is special or starts with a vowel:
func (line *Line) AToAn() {
	str := ""
	slc := []string{}
	for _, v := range line.String {
		if v == 32 {
			if str != "" {
				slc = append(slc, str)
				str = ""
			}
			continue
		}
		str += string(v)
	}
	if str != "" {
		slc = append(slc, str)
		str = ""
	}
	fmt.Println(slc)
	for i := 0; i < len(slc)-1; i++ {
		fmt.Println(slc[i])
		if slc[i] == "a" || slc[i] == "A" {
			if IsVowel(rune(slc[i+1][0])) || IsSpecial(slc[i+1]) {
				if slc[i] == "a" {
					slc[i] = "an"
				} else {
					slc[i] = "An"
				}
			}
		}
	}
	if len(slc) == 0 {
		return
	}

	result := new(string)
	*result = slc[0]
	for i, word := range slc {
		if i != 0 {
			*result += " " + word
		}
	}
	line.String = *result
}

// Add a space where ever a punctuation mark is followed by a letter:
func (line *Line) AddSpace() {
	str := ""
	for i, char := range line.String {
		if i < len(line.String)-1 && isPunctuationMark(char) && isLetter(rune(line.String[i+1])) {
			str += string(char) + " "
			continue
		}
		str += string(char)
	}
	line.String = str
}

// The generator fanction to apply all the necessary changes on the text:
func (text *Text) RockAndRoll() {
	temp := text.First
	for temp != nil {
		if temp.String != "" {
			temp.TrimExcessSpaces()
			temp.organizeQuotes()
			temp.HandleFlags()
			temp.handlePunctution()
			// temp.organizeQuotes()
			temp.AToAn()
			temp.HandleFlags()
			temp.AddSpace()
			temp = temp.Next
		} else {
			fmt.Println("Not enough arguments!!!")
		}
	}
}

// Display the Content of the list:
func (line *Line) Display() {
	if line == nil {
		return
	}
	fmt.Println(line.String)
	line.Next.Display()
}
