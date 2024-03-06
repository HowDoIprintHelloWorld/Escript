package src

type Object struct {}

type Operator struct {}

type Control_flow struct {}

type Symbol struct {}

type Line struct {
  Line_number int
  Content []Token
}

type Token struct {
  Type string
  Object_value Object
  Operator_value Operator
  Control_flow_value Control_flow
  Symbol_value Symbol
}

