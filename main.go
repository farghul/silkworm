package main

import (
	"fmt"
	"log"
	"os"
)

// Launch the program and execute according to the supplied flag
func main() {
	var flag string = flags()

	switch flag {
	case "-h", "--help":
		help()
	case "-v", "--version":
		build()
	case "--zero":
		clearout(assets + "temp/")
		serialize()
		sifter()
	default:
		alert("Unknown argument(s) -")
	}
}

// Record a message to the log file and duplicate the output to console
func journal(message string) {
	file, err := os.OpenFile(assets+"logs/silkworm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	log.Println(message)
	fmt.Println(message)
}

// Print a colourized error message
func alert(message string) {
	fmt.Println("\n", bgred, message, halt, reset)
	fmt.Println("\n", bgyellow, "Use -h for more detailed help information ")
	fmt.Println(reset)
	os.Exit(0)
}

// Display the build version of the program
func build() {
	fmt.Println("\n", yellow+"Silkworm", green+bv, reset)
}

// Print help information for using the program
func help() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(green, " -v, --version", reset, "	Display Program Version")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  Adding your path to file if necessary, run:")
	fmt.Println(green, "    silkworm")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/farghul/silkworm.git")
	fmt.Println(reset)
}
