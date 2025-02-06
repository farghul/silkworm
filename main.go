package main

import (
	"fmt"
	"os"
)

// Launch the program and execute according to the supplied flag
func main() {
	var flag string = flags()

	switch flag {
	case "-h", "--help":
		help()
	case "-r", "--run":
		clearout(assets + "temp/")
		serialize()
		sifter()
	case "-v", "--version":
		build()
	case "--zero":
		alert("No flag detected -")
	default:
		alert("Unknown argument(s) -")
	}
}

// Provide and highlight an informational message
func inform(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
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
	fmt.Println(yellow, "\nOperational Flags:")
	fmt.Println(green, " -h, --help", reset, "		Help Information")
	fmt.Println(green, " -r, --run", reset, "           Run Program")
	fmt.Println(green, " -v, --version", reset, "	Display Program Version")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  Adding your path to file if necessary, run:")
	fmt.Println(green, "    silkworm")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/farghul/silkworm.git")
	fmt.Println(reset)
}
