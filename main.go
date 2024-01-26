package main

func main() {

	cliFlags := ParseFlags()

	ReadConfig(&cliFlags)

}
