package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	// "strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("comments-raw.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// fmt.Println("Slice len: ", len(b))
	// fmt.Println(b) // print as a byte slice

	str := string(b) // convert content to a 'string'

	re := regexp.MustCompile("<p>[^<]+</p>")
	strArr := re.FindAllString(str, -1)
	fmt.Println(len(strArr))

	var repl1 string
	var repl2 string
	for idx, val := range strArr {
		repl1 = strings.Replace(val, "<p>", "", -1)
		repl2 = strings.Replace(repl1, "</p>", "", -1)
		fmt.Println(idx, repl2)
	}

	// fmt.Println(str)
	// fmt.Println("String lenth: ", len(str))
	// fmt.Println("yo")
}
