// revproxyhashry computes the hash of a password.
package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"flag"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func run() int {
	cost := flag.Int("cost", 10, "cost of the bcrypt hash")
	version := flag.Bool("version", false, "If set, outputs the version and quits")

	flag.Parse()

	if *version {
		fmt.Printf("1.0.0\n")
		return 0
	}

	logErr := log.New(os.Stderr, "", 0)

	if *cost < 4 {
		logErr.Printf("Cost needs to be at least 4, but got: %d\n", *cost)
		return 1
	}

	var passwd []byte
	var err error

	switch flag.NArg() {
	case 0:
		fmt.Printf("Please enter the password: ")
		passwd, err = terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			logErr.Printf("failed to read the password: %s", err.Error())
			return 1
		}
		fmt.Println()
	case 1:
		passwd = []byte(flag.Args()[0])
	default:
		logErr.Printf(
			"Unexpected number of positional command-line arguments. Expected at most 1, but got: %d",
			flag.NArg())
		return 1
	}

	hsh, err := bcrypt.GenerateFromPassword(passwd, *cost)
	if err != nil {
		logErr.Printf("failed to generate the hash: %s", err.Error())
		return 1
	}

	fmt.Println(string(hsh))

	return 0
}

func main() {
	os.Exit(run())
}
