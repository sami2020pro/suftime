package main

import (
	f "flag"
	"fmt"
	io "io/ioutil"
	t "time"

	c "github.com/fatih/color"
)

func main() {
	// Help File Color
	helpFileColor := c.New(c.FgWhite)

	// Read Help File and Version File
	dataHelpFile, err := io.ReadFile("/etc/suftime/help.txt")
	dataVersionFile, err := io.ReadFile("/etc/suftime/version.txt")

	if err != nil {
		fmt.Printf("%s", err)
	}

	// UTC Time
	utc := t.Now().UTC()

	// Local Time
	time := t.Now()

	// UTC Flag
	utcPattern := f.Bool("utc", false, "utc flag")

	// Local Flag
	localPattern := f.Bool("local", false, "local flag")

	// Version Flag
	versionPattern := f.Bool("version", false, "version flag")

	// Parse Flags
	f.Parse()

	// Text Color | CLI
	yellow := c.New(c.FgYellow)
	blueBackground := yellow.Add(c.BgBlue)

	// CLI Program
	if *localPattern {
		blueBackground.Println(" Time Location:", time.Location(), "\n Time:", time.String()[10:19], "     ") // [10:30]
	} else if *utcPattern {
		blueBackground.Println("", utc, "")
	} else if *versionPattern {
		blueBackground.Println(" suftime", string(dataVersionFile), "")
	} else {
		helpFileColor.Println(string(dataHelpFile))
	}

}

// Other Sessions
// fmt.Println(time.Location(), "\n", time.String()[10:30])
// fmt.Println(utc)
