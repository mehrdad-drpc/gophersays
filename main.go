// Gophersays is an adaption of the cowsay command.
// Prints a gopher in ASCII art.
package main

import (
	"fmt"
	"os"
	"strings"
)

var gopher = []string{
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

func main() {
	s := strings.Join(os.Args[1:], " ")
	slen := len(s)

	fmt.Printf(" %s\n/ %s /\n %[1]s\n", strings.Repeat("~", slen+2), s)
	fmt.Printf("%s%s\n", strings.Repeat(" ", slen/3), "###")
	fmt.Printf("%s%s\n", strings.Repeat(" ", (slen/3)+2), "##")
	fmt.Printf("%s%s\n", strings.Repeat(" ", (slen/3)+4), "#")
	fmt.Print(strings.Join(gopher, "\n"))
}
