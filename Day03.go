package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)
    isWeird := N%2==1 //by default, odd numbers are weird
    //fmt.Println(isWeird)
    
    switch{
        case N >= 6 && N <= 20:
            isWeird = !isWeird //in this case, evens are weird.
    }
    if isWeird{
        fmt.Println("Weird")
    }else{
        fmt.Println("Not Weird")
    }
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
