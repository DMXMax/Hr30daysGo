package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    
    // Declare second integer, double, and String variables.
    var i2 uint64
    var d2 float64
    var s2 string
    
    // Read and save an integer, double, and String to your variables.
    
if scanner.Scan(){
    

    i2, _ = strconv.ParseUint(scanner.Text(),10,32) 
    
    fmt.Println(i2+i)
}
scanner.Scan()
d2, _ = strconv.ParseFloat(scanner.Text(),10)
    // Print the sum of both integer variables on a new line.
fmt.Printf("%.1f\n",d2+d)
    // Print the sum of the double variables on a new line.
scanner.Scan()
s2=scanner.Text()
fmt.Printf("%s%s",s,s2)
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    }
