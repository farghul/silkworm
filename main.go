package main

// Launch the program and execute according to the supplied flag
func main() {
	var flag string = flags()

	switch flag {
	case "-c", "--create":
		clearout(jira.Programs + "temp/")
		serialize()
		sifter()
	case "-h", "--help":
		help()
	case "-v", "--version":
		build()
	default:
		alert("Unknown argument(s) supplied -")
	}
}
