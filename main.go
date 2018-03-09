package main

import "fmt"
import "math/big"
import "math/rand"


var numbers = "0123456789"

func srand(size int) string {
    buf := make([]byte, size)
    for i := 0; i < size; i++ {
        buf[i] = numbers[rand.Intn(len(numbers))]
    }
    return string(buf)
}


func main() {
  // First number we're analysing
  i := big.NewInt(1)
  // Collatz stack status
  n := big.NewInt(0)

  // Somewhere to store the mod for even/odd checks
  mod := big.NewInt(0)

  // 1 = one , 2 = two, 3 = three ( I swear )
  one := big.NewInt(1)
  two := big.NewInt(2)
  three := big.NewInt(3)

  // Loop forever
  for {
   // Generate a random number with lenght 5000 in base 10
   // Comment if you want to switch to sequential mode ( and uncomment last line of this loop )
   i.SetString(srand(5000), 10)

   // Set our collatz stack with this number
   n.Set(i)
   // Some info
   fmt.Println("Starting with" , i.String())

   // while number is bigger than 1 then
   for n.Cmp( one ) > 0 {
    // Check if it's even
    if ( mod.Mod(n, two).Cmp(one) == 0 ) {
     // If odd multiplie by 3
     n.Mul(n, three)
     // And add one
     n.Add(n, one)
    } else {
     // Otherwise divide by 2
     n.Div(n, two)
    }
   }

  // We're done !
  fmt.Println("Done with" , i.String())

  // Uncomment to switch to sequential mode
  //i.Add(i, one)
  }
}
