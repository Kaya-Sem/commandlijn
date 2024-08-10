package main

import "github.com/Kaya-Sem/commandlijn/cmd"

const Version = "0.0.0"

func main() {
	cmd.Version = Version
	cmd.Execute()
}
