# CPP image for Sandbox

## Initial setup

Refer to the main [README.md](../../README.md) for the setting up the environment.

## Building the image

The image can be built by running the following command. The path to the code to be executed needs to be passed as a build argument.

```
docker build --build-arg PATH_TO_CODE=$PATH_TO_CODE -t cpjudge/cpp .
```

## Running the container

The testcases directory needs to be mounted onto the container. The following command can be executed to run the container:

```
docker run -v $PATH_TO_TESTCASES:/sandbox/testcases cpjudge/cpp
```

## Limiting system resources

Additional options with the `run` command for limiting resources can be used in the following way:

**Memory Limit**

```
docker run -v $PATH_TO_TESTCASES:/sandbox/testcases --memory=<memory_limit>[b|k|m|g] cpjudge/cpp
```

The suffix of `b`, `k`, `m`, `g` indicate bytes, kilobytes, megabytes, or gigabytes.

**CPU Limit**

For Docker 1.13 and higher:

```
docker run -v $PATH_TO_TESTCASES:/sandbox/testcases --cpus=<cpu_limit> cpjudge/cpp
```

For Docker 1.12 and lower:

```
docker run -v $PATH_TO_TESTCASES:/sandbox/testcases --cpu-period=<cpu_period> --cpu-quota=<cpu_quota> cpjudge/cpp
```

Refer to [Limit a container's resources](https://docs.docker.com/config/containers/resource_constraints/) for more details on how to set the limits.
