package main

import (
  "fmt"
  "time"
  "math/rand"
)

func f(n int) {

fmt.Println(n); 

for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }

   fmt.Println("millisecond:",time.Millisecond);
   time.Sleep(time.Millisecond * 1000)
   
}

func main() {
  for i := 0; i < 10; i++ {
    go f(i)
    //f(i)
  }

  fmt.Println("===without using go routine==========");

  for i := 0; i < 10; i++ {
   // go f(i)
   f(i)
  }
  
  fmt.Println("===end without using go routine==========");

  var input string
  fmt.Scanln(&input)

}



