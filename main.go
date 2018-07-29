package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.Join(os.Args[1:], " ")

	fmt.Printf(" %s\n/ %s /\n %[1]s\n", strings.Repeat("-", len(s)+2), s)
}
