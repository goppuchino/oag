package main

import (
	"fmt"
	"github.com/goppuchino/oag"
	"github.com/goppuchino/oag/pkg/file"
	"github.com/goppuchino/oag/pkg/generator"
	"time"
)

func main() {
	fmt.Println(" _______________________")
	fmt.Println(" __  __ \\__    |_  ____/")
	fmt.Println(" _  / / /_  /| |  / __  ")
	fmt.Println(" / /_/ /_  ___ / /_/ /  ")
	fmt.Println(" \\____/ /_/  |_\\____/   Version:", oag.Version, "")
	fmt.Println()
	root := "./"

	start := time.Now()
	spec, err := generator.GenerateOpenAPISpec(root)
	duration := time.Since(start)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = file.SaveAsJSON(spec, "docs/openapi.json")
	if err != nil {
		fmt.Println("Error create JSON file:", err)
		return
	}

	err = file.SaveAsYAML(spec, "docs/openapi.yaml")
	if err != nil {
		fmt.Println("Error create YAML file:", err)
		return
	}

	fmt.Println("OpenAPI specification successfully generated in: ", duration, "")
}
