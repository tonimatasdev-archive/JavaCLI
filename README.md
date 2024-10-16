# JavaCLI

It is a simple program to run specific java version.

## Installation

1. [Download](https://github.com/TonimatasDEV/JavaCLI/releases/tag/1.0.0)
   or [compile](https://github.com/TonimatasDEV/JavaCLI/tree/master?tab=readme-ov-file#how-to-compile) this repo.
2. Put it anywhere and add it to `Environment Variables`.

## Commands

#### javacli

- Example: javacli
- Description: Help command

#### javacli -<version> <java-arguments...>

- Example: javacli -8 -jar X.jar
- Description: Execute specific java version command.

#### javacli -set <version> <path>

- Example: javacli -set 8 C:/Users/x/.jdks/temurin8/java.exe
- Description: Set the path of a specific java version.

#### javacli -download <version> (Unfinished)

- Example: javacli -download 8
- Description: Download and set a java version.

## How to compile

1. Install go 1.23.
2. Run `go build` in the source folder.
3. In the same folder is the JavaCLI.exe.