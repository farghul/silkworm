package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Write a passed variable to a named file
func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0666))
}

// Run a terminal command using flags to customize the output
func execute(variation, task string, args ...string) []byte {
	osCmd := exec.Command(task, args...)
	switch variation {
	case "-e":
		exec.Command(task, args...).CombinedOutput()
	case "-c":
		result, err := osCmd.Output()
		inspect(err)
		return result
	case "-v":
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
		err := osCmd.Run()
		inspect(err)
	}
	return nil
}

// Check for errors, halt the program if found
func inspect(err error) {
	if err != nil {
		panic(err)
	}
}

// Read any file and return the contents as a byte variable
func read(file string) []byte {
	outcome, problem := os.ReadFile(file)
	inspect(problem)
	return outcome
}

// Record a list of files in a folder
func ls(folder string) []string {
	var content []string
	dir := expose(folder)

	files, err := dir.ReadDir(0)
	inspect(err)

	for _, f := range files {
		content = append(content, f.Name())
	}
	return content
}

// Open a file for reading and return an os.File variable
func expose(file string) *os.File {
	outcome, err := os.Open(file)
	inspect(err)
	return outcome
}

// Empty the contents a folder
func clearout(path string) {
	list := ls(path)
	for _, file := range list {
		sweep(path + file)
	}
}

// Remove files or directories
func sweep(cut ...string) {
	inspect(os.Remove(cut[0.]))
}

// Record a message to the log file
func journal(message string) {
	file, err := os.OpenFile(jira.Path+"logs/silkworm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	log.Println(message)
}

// Print a colourized error message
func alert(message string) {
	fmt.Printf("\n%s %s\n", message, halt)
}

// Display the build version of the program
func build() {
	fmt.Println("\nSilkworm", bv)
}

// Print help information for using the program
func about() {
	fmt.Println("\nUsage:")
	fmt.Println("  [program] [flag]")
	fmt.Println("\nExample:")
	fmt.Println("  Adding your path to file if necessary, run:")
	fmt.Println("    silkworm")
	fmt.Println("\nAdditional Options:")
	fmt.Println("  -h, --help", "		Help Information")
	fmt.Println("  -v, --version", "	Display App Version")
	fmt.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	fmt.Println("    https://github.com/farghul/silkworm.git")
	fmt.Println()
}
