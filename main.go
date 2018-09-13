// Gophersays is an adaption of the cowsay command.
// Prints a gopher in ASCII art.
package main

import (
	"flag"
	"fmt"
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
	scream := flag.Bool("scream", false, "makes the gopher scream")

	flag.Parse()
	m := strings.Join(flag.Args(), " ")
	ml := len(m)

	switch {
	case ml <= 0:
		m = "What am I supposed to say?"
		if *scream {
			m = strings.ToUpper(m)
		}
		ml = len(m)
		break
	case ml >= 1 && *scream:
		m = strings.ToUpper(m)
	}

	fmt.Printf(" %s\n/ %s /\n %[1]s\n", strings.Repeat("~", ml+2), m)
	fmt.Printf("%s%s\n", strings.Repeat(" ", ml/3), "###")
	fmt.Printf("%s%s\n", strings.Repeat(" ", (ml/3)+2), "##")
	fmt.Printf("%s%s\n", strings.Repeat(" ", (ml/3)+4), "#")
	fmt.Print(strings.Join(gopher, "\n"))
}
