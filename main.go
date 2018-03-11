package main

import "fmt"
import "math/big"
import "math/rand"
import "os"
import "time"
import "strconv"

var figures = "0123456789"

func isOdd(number string) bool {
  last_val := number[len(number)-1:]
  if ( last_val == "0" || last_val == "2" || last_val == "4" || last_val == "6" || last_val == "8" ) {
    return false
  }
  return true
}

func srand(size int) string {
    rand.Seed( time.Now().UnixNano())
    buf := make([]byte, size)
    for i := 0; i < size; i++ {
        buf[i] = figures[rand.Intn(len(figures))]
    }
    return string(buf)
}

func main() {
  // Check arguments
  if ( len(os.Args) < 2 ) {
    fmt.Println("Must specify the size of the number in figures")
    os.Exit(1)
  }
  figures_count, err := strconv.Atoi(os.Args[1])
  if ( err != nil ) {
    fmt.Println(os.Args[1], "is not an integer")
    os.Exit(2)
  }

  // First number we're analysing
  i := big.NewInt(1)
  // Collatz stack status
  n := big.NewInt(0)

  // 1 = one , 2 = two, 3 = three ( I swear )
  one := big.NewInt(1)
  two := big.NewInt(2)
  three := big.NewInt(3)

  // Init some counters
  var steps int64 = 0
  var previous_time int64 = 0
  var biggest_steps int64 = 0

  // Start from a random number with xxx figures
  i.SetString(srand(figures_count), 10)  

  // Loop forever
  for {
   // Init our collatz stack
   n.Set(i)
   // Reset counter steps
   steps = 0

   // Some info
   // while number is bigger or equal than stopping time
   for n.Cmp( i ) != -1 {
    // Check if it's even
    if ( isOdd(n.String()) ) {
     // If odd multiplie by 3
     n.Mul(n, three)
     // And add one
     n.Add(n, one)
    } else {
     // Otherwise divide by 2
     n.Div(n, two)
    }
    // Count a step
    steps = steps + 1

    // Display status every seconds
    if( previous_time != time.Now().Unix()) {
      // Display status
      fmt.Println(i.String(), "is at", steps, "steps")
      // Save time counter
      previous_time = time.Now().Unix()
    }
   }
   if ( biggest_steps < steps ) {
     fmt.Println("Biggest steps : ", i.String(), "has", biggest_steps, "steps")
     biggest_steps = steps
   }

  // We're done for this number, add 1 and go again
  i.Add(i, one)
  }
}
