package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// define build matrix
	oses := []string{"linux"}
	arches := []string{"amd64"}
	goVersions := []string{"1.24"}
	//oses := []string{"linux", "darwin"}
	//arches := []string{"amd64", "arm64"}
	//goVersions := []string{"1.21", "1.22"}

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory(".")

	// create empty directory to put build outputs
	outputs := client.Directory()

	for _, version := range goVersions {
		// get `golang` image for specified Go version
		imageTag := fmt.Sprintf("golang:%s", version)
		golang := client.Container().From(imageTag)
		// mount cloned repository into `golang` image
		golang = golang.WithDirectory("/src", src).WithWorkdir("/src")

		for _, goos := range oses {
			for _, goarch := range arches {
				// create a directory for each os, arch and version
				path := fmt.Sprintf("build/%s/%s/%s/", version, goos, goarch)
				// set GOARCH and GOOS in the build environment
				build := golang.WithEnvVariable("GOOS", goos)
				build = build.WithEnvVariable("GOARCH", goarch)

				// build application
				build = build.WithExec([]string{"go", "build", "-o", path})

				// get reference to build output directory in container
				outputs = outputs.WithDirectory(path, build.Directory(path))
			}
		}
	}
	// write build artifacts to host
	_, err = outputs.Export(ctx, ".")
	if err != nil {
		return err
	}
	return nil
}
