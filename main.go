package main

import (
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

const defaultDockerAPIVersion = "v1.39"

func main() {
	language := "python"
	// codePath := "code.py"
	questionID := "1"
	testcasesPath := "/media/pvgupta24/MyZone/Projects/go/src/github.com/cpjudge/sandbox/testcases/" + questionID
	submissionPath := "/media/pvgupta24/MyZone/Projects/go/src/github.com/cpjudge/sandbox/temp/"

	go run(testcasesPath, submissionPath, language)
}

func run(testcasesPath string, submissionPath string, language string) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(defaultDockerAPIVersion))
	if err != nil {
		panic(err)
	}
	submissionID := uuid.New().String()
	submissionDirectory := submissionPath + submissionID
	createDirIfNotExist(submissionDirectory)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "cpjudge/" + language,
		Tty:   true,
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: testcasesPath,
				Target: "/sandbox/testcases",
			},
			{
				Type:   mount.TypeBind,
				Source: submissionDirectory,
				Target: "/sandbox/submission",
			},
		},
	}, nil, submissionID)
	fmt.Println(resp.ID)
	fmt.Println(submissionID)

	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
