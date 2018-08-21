package main
import (
    "fmt"
    "bufio"
    "os"
    "io"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    fmt.Println("Hello, World.")
    reader := bufio.NewReader(os.Stdin)
    
    text, err := reader.ReadString('\n')
    if( err != nil  && err != io.EOF ){
        fmt.Println(err)
    }else{
    fmt.Println(text)
    }
    
        
}
