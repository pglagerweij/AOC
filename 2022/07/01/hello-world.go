package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var inputFile string = "trial.txt"

type Directory struct {
	name    string
	subDirs []string
	files   []map[string]string
}

func main() {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Could not read the file due to this %s error \n", err)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	totalInput := strings.Split(strings.ReplaceAll(fileContent, "\n$ ", ";"), ";")
	allFiles := map[string]Directory{}
	currentPath := "/"
	// print file content
	// fmt.Println(fileContent)
	for _, element := range totalInput {
		// fmt.Printf("%v\n", element)
		if strings.HasPrefix(element, "cd ") {
			// fmt.Printf("%v\n", element)
			currentPath = updatePath(currentPath, element)
			// fmt.Println(currentPath)
		} else if strings.HasPrefix(element, "ls") {
			// fmt.Printf("next ls command: %v\n", element)
			allFiles = fillDirectory(allFiles, element, currentPath)
		}
	}
	fmt.Printf("%v", allFiles)
}

func updatePath(currentPath string, commandLine string) string {
	command := strings.Split(commandLine, " ")[1]
	if command == ".." {
		i := len(currentPath) - 2
		for os.IsPathSeparator(currentPath[i]) {
			i--
		}
		updatedPath := currentPath[:i]
		return updatedPath
	} else {
		updatedPath := currentPath + command + "/"
		return updatedPath
	}

}

func fillDirectory(currentDict map[string]Directory, commandLine string, path string) map[string]Directory {
	allCommands := strings.Split(commandLine, "\n")
	allDirs := []string{}
	allFiles := []map[string]string{}
	for i := 1; i < len(allCommands); i++ {
		// fmt.Printf("first thing is %v\n", allCommands[i])
		if strings.HasPrefix(allCommands[i], "dir ") {
			// append to directory
			directory := strings.Split(allCommands[i], " ")[1]
			allDirs = append(allDirs, directory)
		} else {
			fileInformation := strings.Split(allCommands[i], " ")
			fileInfo := make(map[string]string)
			fileInfo["fileSize"] = fileInformation[0]
			fileInfo["fileName"] = fileInformation[1]
			allFiles = append(allFiles, fileInfo)
		}
	}
	currentDir := Directory{}
	currentDir.name = path
	currentDir.files = allFiles
	currentDir.subDirs = allDirs
	currentDict[path] = currentDir
	return currentDict
}
