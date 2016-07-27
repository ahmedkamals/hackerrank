package compression

import (
	"strconv"
)

type StringUtil struct {

}

func NewStringUtil() *StringUtil{
	return new (StringUtil)
}

func (self *StringUtil) Compress(str string) string{

	length := len(str)

	switch length {
	case 0:
		return ""
	case 1:
		return str

	}

	skippable := map[string]bool {
		"p": true,
		"q": true,
		"r": true,
		"s": true,
		"t": true,
		"u": true,
	}

	char := ""
	previousChar := ""
	var output = ""
	var charCount = 0
	endOfString := false
	for index, value := range str {

		char = string(value)
		charCount++

		if(length - 1 == index) {
			endOfString = true
			if (char == previousChar && !skippable[previousChar]) {
				charCount++
			}
		} else if skippable[previousChar] {
			output += previousChar
			previousChar = char
			charCount = 0
			continue
		}

		if(previousChar != char || endOfString) {

			output += previousChar

			if(1 < charCount) {
				output += strconv.Itoa(charCount)
			}

			if(endOfString && (char != previousChar || skippable[previousChar])) {

				output += char
			}

			charCount = 0
		}

		previousChar = char
	}

	return output
}
