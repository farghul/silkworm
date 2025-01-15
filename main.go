package main

// Launch the program and execute according to the supplied flag
func main() {
	var flag string = flags()

	switch flag {
	case "-h", "--help":
		help()
	case "-v", "--version":
		build()
	case "--zero":
		clearout(location + "temp/")
		serialize()
		sifter()
	default:
		alert("Unknown argument(s) -")
	}
}
