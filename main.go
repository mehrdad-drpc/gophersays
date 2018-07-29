package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	s := strings.Join(os.Args[1:], " ")
	d, err := ioutil.ReadFile("./gopher-ascii.txt")

	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf(" %s\n/ %s /\n %[1]s\n", strings.Repeat("~", len(s)+2), s)
	fmt.Printf("%s%s\n", strings.Repeat(" ", len(s)/3), "###")
	fmt.Printf("%s%s\n", strings.Repeat(" ", (len(s)/3)+2), "##")
	fmt.Printf("%s%s\n", strings.Repeat(" ", (len(s)/3)+4), "#")
	fmt.Print(string(d))
}
