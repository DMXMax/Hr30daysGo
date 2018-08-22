package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)
func splitString(str string){
    var (         
        c []byte = make([]byte, len(str), len(str))
        v []byte = make([]byte, len(str), len(str))
        )
    cc := 0
    vv := 0
    
    for z,char := range([]byte(str)){
        if z %2 == 1 {
            v[vv]= char
            vv = vv+1
        } else {
            c[cc] = char
            cc=cc+1
        }
        
    }    

    together := append(c," "...)
    together = append(together, v...)

    fmt.Println(strings.TrimSpace(string(together)))
    
    
    return 
}
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var i int
    var err error
    
    reader := bufio.NewReader(os.Stdin)
    test, _ := reader.ReadString('\n')
    i,err= strconv.Atoi(strings.TrimSuffix(test,"\n"))
    if err != nil {
        fmt.Println(err)
    } else { 
    
        for x :=0; x < i; x++{
            buf, _ := reader.ReadString('\n')
            splitString(strings.TrimSuffix(buf, "\n"))
        }
    }
    

}
