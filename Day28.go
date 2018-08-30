package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Name  string
	Email string
}

func getGoogleEntries(ar []Entry) []Entry {
	var result []Entry
	re := regexp.MustCompile("@gmail.com$")
	for _, val := range ar {
		if re.MatchString(val.Email) {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	arEmails := make([]Entry, N, N)

	for NItr := 0; NItr < int(N); NItr++ {
		firstNameEmailID := strings.Split(readLine(reader), " ")

		firstName := firstNameEmailID[0]

		emailID := firstNameEmailID[1]

		arEmails[NItr] = Entry{firstName, emailID}
	}

	//fmt.Println(arEmails)
	arFiltered := getGoogleEntries(arEmails)

	sort.Slice(arFiltered, func(i, j int) bool {
		return arFiltered[i].Name < arFiltered[j].Name
	})

	for _, val := range arFiltered {
		fmt.Println(val.Name)
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
