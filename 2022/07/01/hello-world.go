package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

var inputFile string = "input.txt"

type Directory struct {
	name    string
	subDirs []string
	files   []map[string]string
	dirSize int
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
	// fmt.Printf("%v", allFiles)
	maxdepth := 0
	allDepths := map[int][]string{}
	for k := range allFiles {
		directoryDepth := strings.Count(k, string('/'))
		allDepths[directoryDepth] = append(allDepths[directoryDepth], k)
		if directoryDepth > maxdepth {
			maxdepth = directoryDepth
		}
		// fmt.Printf("dir %v has depth %v \n", k, directoryDepth)
	}
	// fmt.Printf("%v", allDepths)
	// fmt.Printf("%v", maxdepth)
	currentDepth := maxdepth
	for currentDepth > 0{
		allFiles = calculateSizes(allFiles, allDepths, currentDepth)
		currentDepth--
	}

	
	totalResult := 0
	for k := range allFiles {
		currentDirSize := allFiles[k].dirSize
		if currentDirSize <= 100000 {
			// fmt.Printf("current dir %v has size %v\n", k, currentDirSize)
			totalResult += currentDirSize
			// fmt.Printf("Total is %v\n", totalResult)
		}
		
	}
	// fmt.Printf("%v main dir total is\n", allFiles)
	// fmt.Printf("/fml/bztjtqg/dwhl/wlfprc/ total is %v\n", allFiles["/fml/bztjtqg/dwhl/wlfprc/"].dirSize)
	// fmt.Printf("/fml/bztjtqg/dwhl/fqdzv/ total is %v\n", allFiles["/fml/bztjtqg/dwhl/fqdzv/"].dirSize)
	fmt.Printf("/fml/bztjtqg/dwhl/ total is %v\n", allFiles["/"].dirSize)
	fmt.Printf("The total weigth is: %v", totalResult)

}

func calculateSizes(allFiles map[string]Directory, allDepths map[int][]string,  current int) (map[string]Directory) {
	dirsTocheck := allDepths[current]
	for _, dir := range dirsTocheck {
		// fmt.Printf("dir we are checking: %v\n", dir)

		// fmt.Printf("all files in current dir: %v\n", allFiles[dir].files)
		totalDirSizeFiles := 0
		for _, file := range allFiles[dir].files {
			// fmt.Printf("file we are checking: %v\n", file)
			// fmt.Printf("filesize we are checking: %v\n", file["fileSize"])
			currentSize, _ := strconv.Atoi(file["fileSize"])
			totalDirSizeFiles += currentSize
			
		}
		// fmt.Printf("total files size: %v\n", totalDirSizeFiles)

		// fmt.Printf("all firs in current dir: %v\n", allFiles[dir].subDirs)
		totalDirSizeDirs := 0
		for _, dirsToLoop := range allFiles[dir].subDirs {
			// fmt.Printf("dirs we are checking: %v\n", dirsToLoop)
			// fmt.Printf("%v\n", dir + dirsToLoop + "/")
			// fmt.Printf("%v\n", allFiles[dir + dirsToLoop + "/"].dirSize)
			currentSizeDir := allFiles[dir + dirsToLoop + "/"].dirSize
			// fmt.Printf("%v\n", currentSizeDir)
			// currentSizeDir := 0
			totalDirSizeDirs += currentSizeDir
			
		}
		// allFiles[dir].dirSize = totalDirSizeFiles
		// fmt.Printf("total dir size: %v\n", totalDirSizeDirs)
		currentDir := allFiles[dir]
		currentDir.dirSize = totalDirSizeFiles + totalDirSizeDirs
		allFiles[dir] = currentDir
	}
	return allFiles
}

func updatePath(currentPath string, commandLine string) string {
	command := strings.Split(commandLine, " ")[1]
	if command == ".." {
		i := len(currentPath) - 2
		for os.IsPathSeparator(currentPath[i]) == false {
			i--
		}
		updatedPath := currentPath[:i] + "/"
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
