package main
import (
    "fmt"
    "bufio"
    "strings"
    "io"
    "strconv"
    "os"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
  
   
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
  
    ar:= make([]int, nCount)
  
    for i,num := range strings.Split(readLine(reader)," "){
      ar[i],_ = strconv.Atoi(num)
    }

    totalSwap := 0
  
    for y := 0; y < len(ar); y++{
        swap := 0
        for z:=0; z < len(ar)-1; z++{
            if ar[z] > ar[z+1]{
                ar[z], ar[z+1]= ar[z+1], ar[z]
                swap++
                totalSwap++
            }
        }
        if swap ==0{
            break
            }
        }
  

    fmt.Println("Array is sorted in", totalSwap, "swaps.")
    fmt.Println("First Element:", ar[0])
    fmt.Println("Last Element:", ar[len(ar)-1])

}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }

}
