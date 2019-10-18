// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// This program generates the version and version hash. It is invoked by running `go generate` from the root command
package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	// "time"

	"text/template"
)

const versionTemplate = `// Code generated by go generate; DO NOT EDIT.
// This file was generated by 🤖
package version

const Version = "{{.Version}}"
var GitHash string
var Platform string
`

type versionInfo struct {
	// Timestamp time.Time
	Version string
}

func main() {
	path, _ := os.Getwd()
	dest := filepath.Join(path, "internal", "pkg", "version", "version.go")

	f, err := os.Create(dest)
	if err != nil {
		log.Fatalf("Unable to create output version file: %v", err)
	}
	defer f.Close()

	t := template.Must(template.New("").Parse(versionTemplate))
	versionFile := filepath.Join(path, "VERSION")
	data, err := ioutil.ReadFile(versionFile)
	if err != nil {
		log.Fatalf("Unable to read file version: %v", err)
	}
	version := strings.TrimSpace(string(data))

	info := versionInfo{
		Version: version,
	}

	err = t.Execute(f, info)
	if err != nil {
		log.Fatalf("Error applying template: %v", err)
	}
}
