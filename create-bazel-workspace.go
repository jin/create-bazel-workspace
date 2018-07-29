// Copyright 2018 The Bazel Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// limitations under the License.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
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
	err := os.Mkdir(*outputDir, os.ModePerm)
	panicIf(err)

	// TODO: refactor file logic
	workspaceFile, err := os.OpenFile(*outputDir+"/WORKSPACE", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	panicIf(err)
	defer workspaceFile.Close()

	buildFile, err := os.OpenFile(*outputDir+"/BUILD.bazel", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	panicIf(err)
	defer buildFile.Close()

	instructionsFile, err := os.OpenFile(*outputDir+"/instructions.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	panicIf(err)
	defer buildFile.Close()

	initializeLayer("base", workspaceFile, buildFile, instructionsFile)
	for _, layer := range layers {
		initializeLayer(layer, workspaceFile, buildFile, instructionsFile)
		walkExamplesDirectory(layer, *outputDir)
	}

	log.Println("Workspace successfully created in " + *outputDir + "/")
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

func walkExamplesDirectory(layer string, outputDir string) {
	box, err := rice.FindBox(layer)
	panicIf(err)

	// e.g. <root>/examples/android
	layerExamplesDir := filepath.Join(outputDir, "examples", layer)
	os.MkdirAll(layerExamplesDir, os.ModePerm)

	err = box.Walk("examples", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", ".", err)
			return err
		}
		if info.IsDir() {
			if path != "" {
				os.MkdirAll(filepath.Join(layerExamplesDir,
					strings.TrimPrefix(path, "examples/")), os.ModePerm)
			}
		} else {
			outputFile, err :=
				os.Create(filepath.Join(layerExamplesDir,
					strings.TrimPrefix(path, "examples/")))
			panicIf(err)
			defer outputFile.Close()
			writeToFile(outputFile, readFileContent(filepath.Join(layer, path)))
		}
		return nil
	})

	if err != nil {
		// TODO: Enforce examples
		// fmt.Printf("error walking the path %q: %v\n", ".", err)
	}
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
