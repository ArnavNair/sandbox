# Sandbox using Docker

This repository contains images for running codes in different languages against certain testcases to effectively sandbox the execution.

## Initial Setup

* Install the latest `stable` version of [Docker](https://docs.docker.com/install/).
* Set the path to the code to be executed in an environment variable
  ```
  export PATH_TO_CODE="<path_to_code>"
  ```
* Save each testcase in a separate file, all of which are present in one directory. Make sure that only the testcases for this code are present in that directory.
* Set the path to this directory containing all testcases in an environment variable
  ```
  export PATH_TO_TESTCASES="<path_to_testcases_directory>"
  ```

## Build and run the containers

There are separate configurations for each supported language in the `docker` directory. The instructions present in the respective README can be followed to build and run the containers.

## Building images for other languages

To support other languages follow these steps:

1. Create a `Dockerfile`. Replace `<image>` with __iron/LANGUAGE:dev__. <br/>
    (Example - `iron/gcc:dev` or `iron/java:1.8-dev`)
    ``` docker    
    FROM <image>

    RUN apk add --no-cache bash

    WORKDIR /sandbox
    COPY docker-entrypoint.sh .
    ENTRYPOINT ["bash", "docker-entrypoint.sh"]
    ```
    Refer [iron-io](https://github.com/iron-io/dockers) for more details.

2. Now, create a `docker-entrypoint.sh` with the following content.<br/>
    (Replace `<extension>` with the appropriate extension for the language.)
    ``` bash    
    #!/bin/bash

    submission_directory="/sandbox/submission"
    testcases_directory="/sandbox/testcases"
    mkdir -p $submission_directory/output $submission_directory/error

    cd $submission_directory
    mv code code.<extension>
    ```
3. Once this is done write further code to compile the program `code.<extension>` and redirect output as follows: <br/>
    (Redirection here handles compilation output.)
    ``` bash
    <Bash command for compiling "code.<extension>"> 2>"$submission_directory/compilation_error.err"
    if [[ $? -ne 0 ]]; then
        exit 145 # Compilation error. Terminate.
    fi
    ```

4. Once compiled run the program and redirect the output as follows: <br/>
    (Redirection here handles output and runtime errors if any.)

    ``` bash
    for testcase in $(ls $testcases_directory); do
        <code to run the program> 0<"$testcases_directory/$testcase" \
            1>"$submission_directory/output/$testcase" \
            2>"$submission_directory/error/$testcase"
    done
    ```
5. Final file structure should look like this.
        
    * /sandbox <br/>
        + /docker <br/>
            + /\<language\> <br/>
              + /docker-entrypoint.sh <br/>
              +  /Dockerfile <br/>
              + /README.md (_Optional_) <br/>

6. Now to build a image with this docker file run the command from sandbox folder. <br/>
    ``` bash
    docker build -t cpjudge/<language> ./docker/<language>`
    ```
    Note: Here, we explicitly name the image built as `cpjudge/<language>`. <br/>
    This is essential because `sandbox.go` (Sandbox service) expects the image to be named in this format.

## References
* [Time Limit Constraints per language](https://blog.codechef.com/2009/04/01/announcing-time-limits-based-on-programming-language/)