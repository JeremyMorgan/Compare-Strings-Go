package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	sampleString := "Immanuel1234"

	file, err := os.Open("names.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
//		if compareOperators(sampleString, scanner.Text()) {
//		if compareString(sampleString, scanner.Text()) {
		if compareEF(sampleString, scanner.Text()) {
			fmt.Printf("We found a match!")
		}
	}
	file.Close()

}

/*
func compareString(a string, b string) bool {
	if strings.Compare(strings.ToLower(a), strings.ToLower(b)) == 0 {
		return true
	}
	return false
}
*/

func compareEF(a string, b string) bool {
	if strings.EqualFold(a, b) {
		return true
	}
	return false
}


/*
func compareOperators(a string, b string) bool {
// compare with operators
	if strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}
 */

