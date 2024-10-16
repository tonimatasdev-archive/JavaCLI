package src

import (
	"os"
	"os/exec"
	"strings"
)

func ExecuteJava(javaVersion string) {
	javaPath := GetJavaPath(javaVersion)

	if javaPath == "incompatible" {
		println("Java version " + javaVersion + " is incompatible. If it exits create an issue in https://github.com/TonimatasDEV/JavaCLI/issues")
		os.Exit(1)
	}

	if javaPath == "" {
		println("Set the java path with javacli -set <version> <path>")
		os.Exit(1)
	}

	args := os.Args[2:]

	println("Java "+javaVersion+" with args: ", strings.Join(args, " "))

	cmd := exec.Command(javaPath, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		println("Error starting Java:", err)
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		println("Error executing Java:", err)
		os.Exit(1)
	}
}
