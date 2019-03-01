package main

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

const defaultDockerAPIVersion = "v1.39"

// func main() {
// 	language := "cpp"
// 	// codePath := "code.py"
// 	questionID := "1"
// 	submissionID := "1"
// 	testcasesPath := "/media/pvgupta24/MyZone/Projects/go/src/github.com/cpjudge/testcases/" + questionID
// 	submissionPath := "/media/pvgupta24/MyZone/Projects/go/src/github.com/cpjudge/submissions/" + submissionID

// 	go run(testcasesPath, submissionPath, language, submissionID)
// }

func RunSandbox(testcasesPath string, submissionPath string, language string, containerName string) int64 {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion(defaultDockerAPIVersion))
	if err != nil {
		panic(err)
	}

	// createDirIfNotExist(submissionPath)

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
				Source: submissionPath,
				Target: "/sandbox/submission",
			},
		},
	}, nil, "cpjudge_"+containerName)
	// fmt.Println(resp.ID)

	if err != nil {
		panic(err)
	}

	// TODO: Check return status of compilation, timelimit
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case status := <-statusCh:
		cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
		fmt.Printf("Status Code %d", status.StatusCode)
		return status.StatusCode
	}
	cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})

	// out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	// out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	// if err != nil {
	// 	panic(err)
	// }

	// io.Copy(os.Stdout, out)

	//TODO: return status code
	return 0
}

// func createDirIfNotExist(dir string) {
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		err = os.MkdirAll(dir, 0755)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
