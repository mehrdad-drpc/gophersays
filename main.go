package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.Join(os.Args[1:], " ")
	sl := len(s)

	talk(s, sl)
}

func talk(s string, sl int) {
	d := strings.Repeat("-", sl+2)

	fmt.Printf(" %s\n/ %s /\n %[1]s\n", d, s)
}
