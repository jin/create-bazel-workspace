package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"
)

func main() {
	outputDir := flag.String("dir", "bazel-workspace", "Set the name of the output directory")

	flag.Parse()

	layers := flag.Args()

	log.Println("Creating Bazel workspace at " + *outputDir + "..")
	os.Mkdir(*outputDir, os.ModePerm)

	workspaceFile, err := os.Create(*outputDir + "/WORKSPACE")
	panicIf(err)
	defer workspaceFile.Close()

	buildFile, err := os.Create(*outputDir + "/BUILD.bazel")
	panicIf(err)
	defer buildFile.Close()

	instructionsFile, err := os.Create(*outputDir + "/instructions.md")
	panicIf(err)
	defer buildFile.Close()

	initializeLayerLoads("base", buildFile)
	for _, layer := range layers {
		initializeLayerLoads(layer, buildFile)
	}

	initializeLayer("base", workspaceFile, buildFile, instructionsFile)
	for _, layer := range layers {
		initializeLayer(layer, workspaceFile, buildFile, instructionsFile)
	}

	log.Println("Workspace successfully created in " + *outputDir + "/")

	box, err := rice.FindBox("android/examples")
	if err != nil {
		panic(err)
	}
	log.Println(box)
}

func initializeLayerLoads(layer string, buildFile *os.File) {
	writeToFile(buildFile, readFileContent(layer+"/loads.bzl")+"\n")
}

func initializeLayer(layer string, workspaceFile *os.File, buildFile *os.File, instructionsFile *os.File) {
	writeToFile(workspaceFile, "### layer dependencies:"+layer+" ###\n\n")
	writeToFile(workspaceFile, readFileContent(layer+"/WORKSPACE.bzl"))
	writeToFile(workspaceFile, "\n")

	writeToFile(buildFile, "# layer targets:"+layer+"\n")
	writeToFile(buildFile, readFileContent(layer+"/BUILD.bazel.bzl"))
	writeToFile(buildFile, "\n")

	postCreateInstructions :=
		"# Instructions for the " +
			strings.TrimSuffix(layer, "/") +
			" layer:\n\n" +
			readFileContent(layer+"/post_create.txt") +
			"\n"
	log.Println(postCreateInstructions)
	writeToFile(instructionsFile, postCreateInstructions)
}

func writeToFile(file *os.File, content string) int {
	bytes, err := file.WriteString(content)
	panicIf(err)
	return bytes
}

func readFileContent(relativePath string) string {
	dir, file := filepath.Split(relativePath)

	box, err := rice.FindBox(dir)
	if err != nil {
		log.Fatal("ERROR: Could not locate layer directory: " + dir)
		panic(err)
	}

	fileContent, err := box.String(file)
	if err != nil {
		// Fail silently
		log.Println("WARNING: No " + file + " in " + dir)
		// log.Fatal(err)
	}

	return fileContent
}

func panicIf(e error) {
	if e != nil {
		panic(e)
	}
}
