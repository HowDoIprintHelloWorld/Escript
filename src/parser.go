package src

import (
	"fmt"
	"os"
	"strings"
)

func get_name_characters() []int {
  var name_characters []int
  // 48-58, 65-91, 97-122, 95
  for i := 0; i < 122; i++ {
    
    name_characters
  } 
  return name_characters
}


func Make_tokens(script string) []Line {
	var lines []Line
	var isString, isLineComment, isFileComment bool
  fmt.Println(isLineComment, isFileComment)
	var stringChar rune
	var isEscaped int
  var name_characters = get_name_characters()
	// operands := []string{"!=", "**", ">=", "=="}
	for line_number, line := range strings.Split(script, "\n") {
		var lastCharacters string
		//var tokens []Token
		for _, character := range line {
			switch character {
			case '\\':
				if isEscaped == 1 {
					lastCharacters += "\\"
				} else {
					isEscaped = 2
				}
			case '"', '\'':
				if isString {
					if character == stringChar {
						isString = false
					} else {
						lastCharacters += string(stringChar)
					}
				} else {
					isString = true
					stringChar = character
				}
			default:
        if int
        fmt.Println(int(character))
				lastCharacters += string(character)
        
			}
		}
		fmt.Println(line_number, line)
	}
	return lines
}

func Read_script(file_name string) string {
	file_content, err := os.ReadFile(file_name)
	if err != nil {
		panic(fmt.Sprintf("File not found: ", err))
	}
	return string(file_content)
}
