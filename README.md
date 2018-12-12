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
