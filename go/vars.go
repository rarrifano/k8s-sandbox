package main

import "fmt"

func main(){
  // Sem tipagem
  var a = "monkey"
  fmt.Println(a)

  // Multiplas vars
  var b, c int = 2, 3
  fmt.Println(b, c)

  // Go infere o tipo
  var d = true
  fmt.Println(d)

  // Caso sem valor, é zero!
  var e int
  fmt.Println(e)

  // A maneira bonita!
  f := "donkey"
  fmt.Println(f)
}
