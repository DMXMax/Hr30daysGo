The Go Playground    Imports 
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type BookDate struct {
	Year  int
	Month int
	Day   int
}

func (b *BookDate) Assign(dateString string) {
	arr := strings.Split(dateString, " ")
	b.Day, _ = strconv.Atoi(arr[0])
	b.Month, _ = strconv.Atoi(arr[1])
	b.Year, _ = strconv.Atoi(arr[2])
}

func calcFine(dateExpected BookDate, dateActual BookDate) int {
	var fine int
	if dateActual.Year > dateExpected.Year {
		fine = 10000
	} else if dateActual.Year == dateExpected.Year &&
		dateActual.Month > dateExpected.Month {
		fine = (dateActual.Month - dateExpected.Month) * 500
	} else if dateActual.Year == dateExpected.Year &&
		dateActual.Month == dateExpected.Month &&
		dateActual.Day > dateExpected.Day {
		fine = (dateActual.Day - dateExpected.Day) * 15
	} else {
		fine = 0
	}

	return fine
}

func readLine(reader *bufio.Reader) (string, error) {
	str, err := reader.ReadString('\n')
	if err == nil || err == io.EOF {
		str = strings.TrimSuffix(str, "\n")
		return str, nil
	} else {
		return "", err
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	daStr, _ := readLine(reader)
	deStr, _ := readLine(reader)
	dateExpected, dateActual := BookDate{}, BookDate{}
	dateExpected.Assign(deStr)
	dateActual.Assign(daStr)
	// fmt.Println(dateOut, dateIn)
	fmt.Println(calcFine(dateExpected, dateActual))
}

