package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isPrime(num int) bool {
	result := true

	if num != 2 {

		if num%2 == 0 || num == 1 {
			result = false
		}
	}

	if result == true { //keep testing

		//count to its square root, and test every odd number since
		var sqCap int = int(math.Sqrt(float64(num)))
		if sqCap > 2 {

			for x := 3; (x <= sqCap) && (result == true); x += 2 {
				if num%x == 0 {
					result = false
				}

			}

		}
	}
	return result

}

func main() {
	var i int
	var err error

	reader := bufio.NewReader(os.Stdin)
	test, _ := reader.ReadString('\n')
	i, err = strconv.Atoi(strings.TrimSuffix(test, "\n"))
	if err != nil {
		fmt.Println(err)
	} else {

		for x := 0; x < i; x++ {
			buf, _ := reader.ReadString('\n')
			buf = strings.TrimSuffix(buf, "\n")
			num, _ := strconv.Atoi(buf)
			if isPrime(num) {
				fmt.Println("Prime")
			} else {
				fmt.Println("Not prime")
			}
		}
	}
}
