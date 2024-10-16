package src

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
)

type Versions struct {
	JAVA8  string `json:"java8"`
	JAVA11 string `json:"java11"`
	JAVA16 string `json:"java16"`
	JAVA17 string `json:"java17"`
	JAVA18 string `json:"java18"`
	JAVA19 string `json:"java19"`
	JAVA20 string `json:"java20"`
	JAVA21 string `json:"java21"`
	JAVA22 string `json:"java22"`
	JAVA23 string `json:"java23"`
}

var VersionsFile string

var versions Versions

func getFilePath() {
	usr, err := user.Current()
	if err != nil {
		println("Error getting the current user directory:", err)
		os.Exit(1)
	}

	folder := filepath.Join(usr.HomeDir, "AppData", "Roaming", "JavaCLI")

	err = os.MkdirAll(folder, 0)

	if err != nil {
		println("Error creating the folder:", err)
		os.Exit(1)
	}

	VersionsFile = filepath.Join(folder, "versions.json")
}

func LoadVersions() {
	getFilePath()
	if _, err := os.Stat(VersionsFile); os.IsNotExist(err) {
		println("Versions file not found, creating it...")

		versions = Versions{
			JAVA8:  "",
			JAVA11: "",
			JAVA16: "",
			JAVA17: "",
			JAVA18: "",
			JAVA19: "",
			JAVA20: "",
			JAVA21: "",
			JAVA22: "",
			JAVA23: "",
		}

		SaveVersions()
		return
	}

	file, err := os.Open(VersionsFile)

	if err != nil {
		println("Error opening the versions file:", err.Error())
		os.Exit(1)
	}

	defer closeFileHandler(file)

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&versions); err != nil {
		println("Error decoding the versions file:", err.Error())
		os.Exit(1)
	}
}

func SaveVersions() {
	var file *os.File
	if _, err := os.Stat(VersionsFile); os.IsNotExist(err) {
		createdFile, err := os.Create(VersionsFile)

		file = createdFile

		if err != nil {
			println("Error creating versions file:", err.Error())
			os.Exit(1)
		}
	} else {
		openedFile, err := os.OpenFile(VersionsFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

		file = openedFile

		if err != nil {
			println("Error opening the versions file:", err.Error())
			os.Exit(1)
		}
	}

	defer closeFileHandler(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(versions); err != nil {
		println("Error encoding versions:", err.Error())
	}
}

func GetJavaPath(version string) string {
	switch version {
	case "8":
		return versions.JAVA8
	case "11":
		return versions.JAVA11
	case "16":
		return versions.JAVA16
	case "17":
		return versions.JAVA17
	case "18":
		return versions.JAVA18
	case "19":
		return versions.JAVA19
	case "20":
		return versions.JAVA20
	case "21":
		return versions.JAVA21
	case "22":
		return versions.JAVA22
	case "23":
		return versions.JAVA23
	default:
		return "incompatible"
	}
}

func SetVersion(version string, path string) {
	switch version {
	case "8":
		versions.JAVA8 = path
	case "11":
		versions.JAVA11 = path
	case "16":
		versions.JAVA16 = path
	case "17":
		versions.JAVA17 = path
	case "18":
		versions.JAVA18 = path
	case "19":
		versions.JAVA19 = path
	case "20":
		versions.JAVA20 = path
	case "21":
		versions.JAVA21 = path
	case "22":
		versions.JAVA22 = path
	case "23":
		versions.JAVA23 = path
	}

	println("Java " + version + " now has the path: " + path)

	SaveVersions()
}

func closeFileHandler(file *os.File) {
	err := file.Close()

	if err != nil {
		println("Error closing the versions file:", err.Error())
	}
}
