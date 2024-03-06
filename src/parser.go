package src

import (
  "os"
  "fmt"
  "strings"
)

func Make_tokens(script string) []Line {
  var lines []Line
  var isString, isLineComment, isFileComment bool
  var stringChar rune;
  var isEscaped int;
  for line_number, line := range strings.Split(script, "\n") {
    var lastCharacters string
    var tokens []Token;
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
            if character == stringChar {isString = false
            } else {lastCharacters += string(stringChar)}
          } else {
            isString = true
            stringChar = character
          }
        case 
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