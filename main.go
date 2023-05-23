// main.go

package main

import (
	"fmt"
	"github.com/git719/utl"
	"os"
	"strings"
)

const (
	prgname = "chk"
	prgver  = "0.5.0"
)

func printUsage() {
	fmt.Printf(prgname + " IP address plus port checker v" + prgver + "\n" +
		"    IP_Address:Port   Check if IP_Address:Port is reachable\n" +
		"    -v                Print this usage page\n")
	os.Exit(0)
}

func main() {
	args := len(os.Args[1:]) // Number of arguments, not including the program itself
	if args != 1 {           // Accept only 1 argument
		printUsage()
	}

	arg1 := os.Args[1]
	if arg1 == "-v" {
		printUsage()
	}

	c := rune(arg1[0])
	if !utl.IsDigit(c) {
		utl.Die("Invalid IP address\n")
	}

	split := strings.Split(arg1, ":")
	if len(split) != 2 || split[1] == "" {
		utl.Die("Invalid port\n")
	}

	if utl.IsIpPortStrReachable(arg1) {
		fmt.Println("Reachable")
	} else {
		fmt.Println("NOT reachable")
	}
}
