package main

import (
	"flag"
	"fmt"
	io "io/ioutil"
	"time"
)

const version string = "v0.2.4"

func main() {
	// Help File Color
	helpFileColor := "\x1b[38;5;15m"

	// Read Help File and Version File
	dataHelpFile, err := io.ReadFile("/etc/suftime/help.txt")

	if err != nil { /* Print logs */
		fmt.Printf("%s", err)
	}

	// UTC Time
	utc := time.Now().UTC()

	// Local Time
	time := time.Now()

	// UTC Flag
	utcPattern := flag.Bool("utc", false, "utc flag")

	// Local Flag
	localPattern := flag.Bool("local", false, "local flag")

	// Version Flag
	versionPattern := flag.Bool("version", false, "version flag")

	// Parse Flags
	flag.Parse()

	// Text Color | CLI
	yellow := "\x1b[38;5;154m"
	blueBackground := "\x1b[48;5;45m"
	reset := "\x1b[0m"

	// CLI Program
	if *localPattern {
		fmt.Print(blueBackground, yellow, " Time Location: ", time.Location(), " ", "\n Time: ", time.String()[10:19], "      \n") // Available: [10:30]
	} else if *utcPattern {
		fmt.Print(blueBackground, yellow, " ", utc, " ", reset, "\n")
	} else if *versionPattern {
		fmt.Print(blueBackground, yellow, " ", "suftime", " ", version, " ", reset, "\n")
	} else {
		fmt.Print(helpFileColor, string(dataHelpFile), reset)
	}
}
