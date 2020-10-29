package main

import (
	f "flag"
	io "io/ioutil"
	t "time"

	c "github.com/fatih/color"
)

func main() {
	// Help File Color
	helpFileColor := c.New(c.FgWhite)

	// Read Help File
	dataHelpFile, _ := io.ReadFile("/etc/suftime/help.txt")

	// UTC Time
	utc := t.Now().UTC()

	// Local Time
	time := t.Now()

	// UTC Flag
	utcPattern := f.Bool("utc", false, "utc flag")

	// Local Flag
	localPattern := f.Bool("local", false, "local flag")

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
	} else {
		helpFileColor.Println(string(dataHelpFile))
	}

}

// Other Sessions
// fmt.Println(time.Location(), "\n", time.String()[10:30])
// fmt.Println(utc)
