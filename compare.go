package main

import "strings"

func main() {

	sampleString := "This is a sample string"
	compareString := "this is a sample string"

	/*
	if strings.Compare(strings.ToLower(sampleString), strings.ToLower(compareString)) == 0 {
		println("The strings match")
	}else {
		println("Strings do not match")
	}
	*/

	/*
	if strings.ToLower(sampleString) == strings.ToLower(compareString) {
		println("The strings match")
	}else {
		println("Strings do not match")
	}
	 */

	if strings.EqualFold(sampleString, compareString) {
		println("The strings match")
	}else {
		println("Strings do not match")
	}


}
