// Gophersays is an adaption of the cowsay command.
// Prints a gopher in ASCII art.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	gopher := []string{
		"                        dhyysoo++++ooossyhd",
		"                Ndyo/:#######################:+sd",
		"      NN/:###ds/#######################/+////+/###+y#######",
		"  Ny/:###/hy:#:+/::###::/+:#########+/:`  .:#  :+/##:oo######",
		" h:##//:+o:#:+:  +hdy.    #+:#####++`     s####` :+###:sN +###:",
		" ###d   /##:o   :#####      o:###//       h#s#N.  :/####+Ny####",
		"d###sN :###s`    `sd+d:     `o###s         ://`    s#####++###/",
		" y:##s/####s                `s###s`               `s######s:/",
		"  Nhso#####//               +:###:o`             `o:######:",
		"    N#######+/            `+/:+sss+o/`         ./+#########",
		"    y########:+/:.    `#://#+N     N:+//:::::/+/###########/",
		"    +############:/+++/:##ooodNNN dsoo+#####################",
		"    /####################s/:::::::::::/s####################h",
		"    :####################+o//+++s/++//oo####################s",
		"    :######################o/  ##  ++/:#####################o",
		"    /######################:+  ::  :/#######################+",
		"    +#######################+//++//+########################+",
	}

	scream := flag.Bool("scream", false, "makes the gopher scream")
	flag.Parse()

	msg := strings.Join(flag.Args(), " ")

	Gopher(&gopher, scream, msg, os.Stdout)
}

func Gopher(ASCIIart *[]string, scream *bool, msg string, out io.Writer) {
	g := strings.Join(*ASCIIart, "\n")
	ml := len(msg)

	if ml <= 0 {
		msg = "What am I supposed to say?"
		ml = len(msg)
	}

	if *scream {
		msg = strings.ToUpper(msg)
	}

	fmt.Printf(" %s\n/ %s /\n %[1]s\n", strings.Repeat("~", ml+2), msg)
	for i := 0; i <= 3; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", (ml/3)+(i*2)), strings.Repeat("#", 3-i))
	}
	io.WriteString(out, g)
}
