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
}

func initializeLayerLoads(layer string, buildFile *os.File) {
	buildFile.WriteString(readFileContent(layer+"/loads.bzl") + "\n")
}

func initializeLayer(layer string, workspaceFile *os.File, buildFile *os.File, instructionsFile *os.File) {
	workspaceFile.WriteString("### layer dependencies:" + layer + " ###\n\n")
	workspaceFile.WriteString(readFileContent(layer + "/WORKSPACE.tpl"))
	workspaceFile.WriteString("\n")

	buildFile.WriteString("# layer targets:" + layer + "\n")
	buildFile.WriteString(readFileContent(layer + "/BUILD.bazel.tpl"))
	buildFile.WriteString("\n")

	post_create_instructions := "# Instructions for the " + strings.TrimSuffix(layer, "/") + " layer:\n\n" + readFileContent(layer+"/post_create.txt") + "\n"
	log.Println(post_create_instructions)
	instructionsFile.WriteString(post_create_instructions)
}

func readFileContent(relative_path string) string {
	dir, file := filepath.Split(relative_path)

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
