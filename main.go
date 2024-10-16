package main

import (
	"JavaCLI/src"
	"os"
	"strings"
)

func main() {
	src.LoadVersions()

	if len(os.Args) < 2 {
		println("JavaCLI 1.0.0 - Commands:")
		println("javacli -<version> <java-arguments>")
		println("javacli -set <version> <path>")
		println("javacli -download <version>")
		return
	}

	firstArgument := strings.Replace(os.Args[1], "-", "", -1)

	if firstArgument == "set" {
		if len(os.Args) == 4 {
			src.SetVersion(os.Args[2], os.Args[3])
		} else {
			println("Usage: javacli -set <version> <path>")
		}

		return
	}

	if firstArgument == "download" {
		return
	}

	src.ExecuteJava(firstArgument)
}
