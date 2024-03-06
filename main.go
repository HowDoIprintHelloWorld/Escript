package main

import (
  "main/src"
)

func main() {
	script_string := src.Read_script("example.es")
  src.Make_tokens(script_string)
}
