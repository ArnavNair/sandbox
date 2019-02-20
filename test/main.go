package main

import (
	"context" // "encoding/base64"
	// "encoding/json"
	"fmt"
	"os"
	// "path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)

const defaultDockerAPIVersion = "v1.39"

func main() {
	Build()
	// cli, err := client.NewClientWithOpts(client.WithVersion(defaultDockerAPIVersion))
	// if err != nil {
	// 	panic(err)
	// }

	// containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	// if err != nil {
	// 	panic(err)
	// }

	// for _, container := range containers {
	// 	fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	// }
}

func buildAPI() error {
	cli, err := client.NewClientWithOpts(client.WithVersion(defaultDockerAPIVersion))
	if err != nil {
		panic(err)
	}
	fmt.Print(cli.ClientVersion())

	// args := map[string]*string{
	// 	"PATH_TO_CODE": "docker/cpp/code.cpp",
	// }

	opt := types.ImageBuildOptions{
		SuppressOutput: false,
		Dockerfile:     "docker/cpp/Dockerfile",
		// BuildArgs:      args,
	}
	_, err = cli.ImageBuild(context.Background(), nil, opt)
	if err == nil {
		fmt.Printf("Error, %v", err)
	}
	return err
}

const archiveLoc = "/media/pvgupta24/MyZone/Projects/go/src/github.com/cpjudge/sandbox/docker/cpp/"

// Build the container using the native docker api
func Build() error {
	dockerBuildContext, err := archive.TarWithOptions(archiveLoc, &archive.TarOptions{})
	defer dockerBuildContext.Close()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithVersion(defaultDockerAPIVersion))

	pathToCode := "code.cpp"
	fmt.Printf("Code: %s\n", pathToCode)
	args := map[string]*string{
		"PATH_TO_CODE": &pathToCode,
	}
	options := types.ImageBuildOptions{
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		//   PullPar/ent:     true,
		//   Tags:           getTags(s),
		//   Dockerfile:     "image/" + s.Os + "/Dockerfile",
		BuildArgs: args,
	}
	buildResponse, err := cli.ImageBuild(context.Background(), dockerBuildContext, options)
	fmt.Println("Hee")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println("Hee2")
	defer buildResponse.Body.Close()
	fmt.Printf("********* %s **********\n", buildResponse.OSType)
	fmt.Println("Hee3")

	termFd, isTerm := term.GetFdInfo(os.Stderr)
	return jsonmessage.DisplayJSONMessagesStream(buildResponse.Body, os.Stderr, termFd, isTerm, nil)
}
