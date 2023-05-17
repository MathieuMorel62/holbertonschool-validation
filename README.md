# <p align=center>Holbertonschool Validation</p>

<img src="https://camo.githubusercontent.com/65f00c4f146b2c20e3c6ae1abc7501a4937e1fb54239fcc41860da0c59d27e56/68747470733a2f2f6d69726f2e6d656469756d2e636f6d2f6d61782f333030302f312a6d717630334b726c47354c4b325855317556344c4a672e676966" width="100%">

## Introduction to DevOps: Automate Everything to Focus on What Really Matters.
## Learning Objectives

This project aims at showing use cases where a DevOps mindset is bringing value to a software project by automating it, which decreases the amount of manual work and increases the development speed. It focuses on why automation is useful and helps speeding a development lifecycle.
  
After this project, you should be able to:
  
- Understand the value of automating tedious tasks
- Define a development lifecycle
- Automate shell-like tasks with Make, and/or shell script
- Be aware of tools dependencies and the value of reproducing environment
- Build static HTML website from Markdown code using Go-Hugo

## Resources
### Read or Watch:
- [Go Hugo Quickstart](https://gohugo.io/getting-started/quick-start/)
- [Git SCM Book](https://git-scm.com/book/en/v2)
- [GNU Make Docs](https://www.gnu.org/software/make/manual/html_node/index.html)
- [Installing & Using Themes](https://intranet.hbtn.io/rltoken/kL8c124W_cnv5apEAdjkhg)
- [Add a Help Target To a Makefile](https://gist.github.com/prwhite/8168133)

## Prerequisites

You should have a basic knowledge of the following concepts:
  
- Shell terminal basics, using command lines:
  
  - Navigating in a Unix file-system
  - Understanding how stdin/stdout redirection and piping
  - Showing and searching the content of a text files
  - Defining and using Environment Variables
  - Adding command lines to your terminal using the `apt-get` package manager and/or with the `PATH` variable
  - Writing and executing a shell script
  
- Git with the command line (and also a graphical interface)

  - Retrieving or creating a repository
  - Manipulating changes locally with Git’s 3 steps process (workspace, staging, history)
  - Distributing changes history with remotes repositories

- Make/Makefile usage:
  
  - Executing tasks through make targets
  - Default target and PHONY target
  - Makefile’s variables and macro syntax

## Tooling
This project needs the following tools / services:
  
- An HTML5-compliant web browser (Firefox, Chrome, Opera, Safari, Edge, etc.)
- A free account on [GitHub](https://intranet.hbtn.io/rltoken/u6680ax-ghu8v-AsFSDbSA), referenced as `<GitHub Handle>`
- A shell terminal with `bash`, `zsh` or `ksh`, including the standard Unix toolset (`ls`, `cd`, etc.) with:
  - GNU Make in version 3.81+
  - Git (command line) in version 2+
  - [Go Hugo](https://gohugo.io) **v0.84.0** (very crucial for the theme that you’ll use).
- The student needs to be able to spawn up a clean Ubuntu 18.04 system. Therefore [Docker](https://www.docker.com) is recommended with NO prior knowledge.
- A text editor or IDE (Integrated Development Editor) of your convenience (Visual Code, Notepad++, Vim, Emacs, IntelliJ, etc.)

## Lifecycle

The lifecycle of this project is organized into several key steps:

1. **Create a new post**: Use the command `make post POST_NAME="post_name" POST_TITLE="Post Title"` to create a new post file in the `content/posts/` directory with a specified name and title.

2. **Edit the content**: Open the created post file and edit its content by adding text, images, and links as needed.

3. **Build the site**: Use the command `make build` to build the website with Hugo and generate the files in the `dist/` directory. This step transforms Markdown files into HTML pages and assembles them into a complete website structure.

4. **Clean**: If you need to remove the `dist/` directory and all the files it contains, use the `make clean` command. This step is useful if you want to rebuild the site from scratch or remove temporary files.

5. **Deployment**: Once you have successfully built the site, you can deploy the generated files in the `dist/` directory to a web server or hosting service of your choice.

6. **Help**: To display the list of available commands and their descriptions, you can use the `make help` command in the terminal. This command extracts the comments associated with each target in the Makefile and displays them to give you a quick overview of the project's features.

# Tasks

### [0. Go-Hugo Website](https://github.com/MathieuMorel62/holbertonschool-validation/tree/main/module1_task0)
For your first day at “Awesome Inc.”, you are tasked to work on the corporate website. The communication department wants you to allow them to write and publish blog posts on the website in total autonomy.
   
Your boss tells you that the website infrastructure has documentation and sends you a link to an internal wiki page that contains the following instructions:

```markdown
# README for "Awesome Inc." Website

Ain't nobody got time to write documentation: I've put all the HTML and CSS into `public/index.html` on the production server.

If you want to update the site: edit the `index.html` file in place so you can go fast.
```
  
  You also receive an email from the communication department:
    
```markdown
Hello dear IT colleague!

Our team is experienced in writing in Markdown language as they use the Go-Hugo static site generator (https://gohugo.io/getting-started/quick-start/) internally.

Do you mind to switch to Go Hugo to avoid writing HTML content for our blog posts?
```
  
Let’s get started: you are expected to provide the source code of a new website that can be built using Go-Hugo.
  
**Requirements**

- Use the theme [“ananke”](https://themes.gohugo.io/themes/gohugo-theme-ananke/) for the website by following this [Quickstart](https://gohugo.io/getting-started/quick-start/#step-3-add-a-theme)
- If you are using the sandboxe, do NOT install `hugo` with the package manager. Download the correct executable or package file from their github.
- Usage of [Git Submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules) is prohibited: there should be no file `.gitmodules`
- The website title should be “Awesome Inc.”
- The contents consists in a single blog post which title should be “Welcome to Awesome Inc.”, stored in a file named `welcome.md`
- All of the website’s source code is stored under a directory named `module1_task0`
- The command line `hugo` in version 0.84.0 must be used
- The website is expected to be generated into the directory `module1_task0/dist/`
- The directory `module1_task0/dist/` must not be committed (it should be **absent** from the repository)

Here is a simple example of the expected generation process:
  
```bash
➜ ls -l dist/ 2>/dev/null | wc -l
0

➜ hugo > /dev/null 2>&1
➜ ls dist/ 2>/dev/null | grep -c 'index.html'
1
```

-------------------------------

### [1. Development Lifecycle With Makefile](https://github.com/MathieuMorel62/holbertonschool-validation/tree/main/module1_task1)
In the DevOps methodology, while tools often change, the “development” lifecycle is generally staying the same. In this project, you will start to define this lifecycle via the following steps:
  
- “build”: Generate the website from the markdown and configuration files in the directory `dist/`.
- “clean”: Cleanup the content of the directory `dist/`
- “post”: Create a new blog post whose filename and title come from the environment variables `POST_TITLE` and `POST_NAME`.
  
Write a Makefile to implement these steps for the actual Go-Hugo website’s source code.
  
**Requirements**
  
- Same requirements as the previous task:
  
  - A Valid Go-Hugo website is provided
  - There are no Git Submodules
  - The theme ananke is installed
  - No directory `dist/` committed
- GNU Make version 3.81 or 4.0 must be used
  
- The “Build” step should be executed by the command `make build`, “Clean” by `make clean` and “Post” by `make post`.

```bash
➜ ls -1 ./dist/

➜ make build
➜ ls -1 ./dist/index.html
index.html

➜ ls -1 ./content/posts/
welcome.md

➜ make POST_NAME=who-are-we POST_TITLE="Who are we" post
➜ ls -1 ./content/posts/
welcome.md who-are-we.md

➜ make clean
➜ ls -1 ./dist/
➜
```

--------------

### [2. Empathy as Code: Inline Help, Comment and Documentation](https://github.com/MathieuMorel62/holbertonschool-validation/tree/main/module1_task2)
As you are thinking about the person who will be in charge of this website after your mission for Awesome Inc. is finished, you want to write proper documentation.
  
You want to write a `README.md` file in [Markdown](https://daringfireball.net/projects/markdown/syntax) syntax which describes the requirements to build the website (Go-Hugo, Make, etc.), describes the lifecycle and its steps.
  
As an accomplished command line user, you also feel that providing a command line “help” could be a great addition, as well as comment for source code readers.
  
**Requirements**
  
- Same requirements as the previous task:

  - A Valid Go-Hugo website is provided
  - There are no Git Submodules
  - The theme `ananke` is installed
  - No directory `dist/` committed
  - `Makefile` present
- Add comments in the Makefile to describe what each target is expected to do.

  - These comments should be written on the same line as the targets
  - Each comment should start with two characters `#`
  
Here is a Makefile example:

```bash
cook:     ## Prepare a meal in the current kitchen
    cook ./kitchen
```

- Add a “help” target to the `Makefile` which prints out the list of targets and their usage.
  - The usage should be the comment associated with each target
  - Do not just `echo` or `print` this help usage:
    - If you ONLY update the comment associated to a target, the “help” target should displayed the updated version.
  
With the same previous example:
  
```bash
➜ make help
cook: Prepare a meal in the current kitchen
help: Show this help usage
```
  
- Add (or update) a README.md file located at the repository’s root
  - A section named “Prerequisites” is expected to describe the requirements for this website (GoHugo, etc.)
  - A section named “Lifecycle” is expected to describe the different steps

```bash
➜ cat README.md | grep -c "## Prerequisites"
1
➜ cat README.md | grep -c "## Lifecycle"
1
```

--------------------------------

### [3. From Development To Production](https://github.com/MathieuMorel62/holbertonschool-validation/tree/main/module1_task3)
Now that you automated the website generation and documented the process, it’s time to deploy this new site to the production environment.
  
Your next goal is to improve the delivery process to help your colleagues from the “production team” (e.g. the team in charge of the production infrastructure). After sending the instructions to the production team, you receive the following answer:
  
```bash
Dear Developer,

I followed the instructions you gave me to generate the website. The command `make build` is failing on the production server with an error. Can you please help me?

## Production Instructions

* The production website is hosted in an Ubuntu 18.04 Docker container
* The applications "GoHugo" and "Make" are installed with `apt-get update && apt-get install -y hugo make`.
* When running the command `make build`, there is a bunch of errors which end with the following lines:

Error: Error building site: logged 15 error(s)
Makefile:2: recipe for target 'build' failed
make: *** [build] Error 255
```
  
So that you understand what is happening to debug the issue, let’s start by reproducing the production environment on your development machine.
  
**Requirements**

- Write a shell script named `setup.sh` to reproduce a pseudo “production” environment locally
  - The script must be executable
  - The script must reproduce all the steps described in the “Production Instructions”
  - The script file is located at the root of your project, along with the `Makefile`
- The script should be executed **in** a disposable Ubuntu 18.04 environment
  - You can spawn a sandboxed environment with the following command (with [Docker](https://intranet.hbtn.io/rltoken/4-LomWsN4dV31c-IwVMwgw)): `docker run --rm --tty --interactive --volume=$(pwd):/app --workdir=/app ubuntu:18.04 /bin/bash` and use the command `exit` to quit (and delete) the environment.
- The result of the script execution should be the same as what is described by your colleague: it must exit with the error code from Make and prints the full error in the stderr.
  
```bash
➜ grep 'UBUNTU_CODENAME' /etc/os-release
UBUNTU_CODENAME=bionic

➜ command -v hugo >/dev/null 2>&1 || echo "No 'hugo'"
No 'hugo'

➜ ./setup.sh >/dev/null 2>&1
➜ command -v hugo >/dev/null 2>&1 || echo "No 'hugo'"
➜ ./setup.sh 2>&1 | grep -c "recipe for target 'build' failed"
1
```

----------------------

### [4. Fixing The Production Deployment](https://github.com/MathieuMorel62/holbertonschool-validation/tree/main/module1_task4)
Now that you are able to reproduce the same error as the production’s, you have to analyse and fix it.  
  
By reading carefully the output from the `setup.sh` script, you are expected to find the error caused by the production’s environment.
  
You have to fix the error by updating the script `setup.sh` to help your colleagues from the production team.
  
**Requirements**
  
- Update the shell script named `setup.sh`
  - The script must be executable
  - The script file is located at the root of your project, along with the `Makefile`
- The script should be executed in a fresh Ubuntu 18.04 environment
  - You can spawn a sandboxed environment with the following command (with [Docker](https://intranet.hbtn.io/rltoken/4-LomWsN4dV31c-IwVMwgw)): `docker run --rm --tty --interactive --volume=$(pwd):/app --workdir=/app ubuntu:18.04 /bin/bash` and use the command `exit` to quit (and delete) the environment.
- The result of the script execution should be successful (with an exit code `0`) with the directory `./dist` containing the generated website.
  
```bash
➜ grep 'UBUNTU_CODENAME' /etc/os-release
UBUNTU_CODENAME=bionic
➜ command -v hugo >/dev/null 2>&1 || echo "No 'hugo'"
No 'hugo'
➜ ls -l dist/ 2>/dev/null | wc -l
0
➜ ./setup.sh >/dev/null 2>&1
➜ echo $?
0
➜ hugo > /dev/null 2>&1
➜ ls dist/ 2>/dev/null | grep -c 'index.html'
1
```

-----------------

# <p align=center>Testing in the Software Development Methodology</p>
<img src="https://www.mindinventory.com/blog/wp-content/uploads/2022/10/golang.gif" width="100%">

## Learning Objectives

This project aims at practicing with automated tests. The goal is to understand the pros and cons of different testing methods to be able to understand the value of doing, or not doing, a kind of test.
  
After this project, you should be able to:
  
- Understand what linting is the extent of its usages (which kind of file can be linted, and the impact of running it often)
- Understand the difference between unit tests and integration tests
- Use code coverage as a helper to write tests
- Understand that not only “classical” code is to be tested, but also a lot of the artifacts we can generate
- Understand how “component”-based testing for acceptance and end to end validation is to be used

## Prerequisites

The following elements are required In addition to the previous module (**“Module 1: Introduction to DevOps: Automate Everything to Focus on What Really Matters”**) prerequisites.
  
## Concepts
You should have a basic knowledge on the following concepts:

- What a compiled language is (C/C#/Golang/Rust/etc.)

  - Generation process from source to executable binary
  - Basic types: string, integer, boolean, maps, arrays
  - Basic algorithmic: loops, conditional, functions
- Installing command line tools with NPM (in addition to package managers)
  
- Understand the basics of the HTTP protocol (client/server, verbs, headers)

## Tooling
This project needs the following tools / services:

- Same tools as previous module
- [Golang](https://go.dev) in `v1.15.*`
- [NPM](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) v7+ with [NodeJS](https://nodejs.org/en/download) `v14.*` (stable)
- [Python 3 with pip module](https://docs.python.org/3/installing/index.html)
- [golangci-lint](https://github.com/golangci/golangci-lint)

## Story

Following the previous module situation, you are now able to build and deploy the static website for the company Awesome Inc. in an automated way.
  
As the communication team is happy with your work, you’ve been tasked to add an HTTP API to improve the website.
  
You want to ensure that the shipped software is without bugs. You’ll test each component of this new website to ensure that there will be no regression in the future, and to make sure that any refactoring or change can be done with confidence.
  
The API is written in the [Golang](https://go.dev) language but there is no need to be familiar with the language.
  
While the production team is building the new production platform for the website, you’ve been tasked to create an HTTP API application to add new features.
  
As we are in a “DevOps” course, your “Ops” personality expects that this application is monitored by a “Health” page to determine if the application is running and ready to accept traffic.
  
## Reference Readings

- [https://tour.golang.org/welcome/1](https://go.dev/tour/welcome/1)
- [http://martinfowler.com/bliki/UnitTest.html](https://martinfowler.com/bliki/UnitTest.html)
- [https://en.wikipedia.org/wiki/Software_testing](https://en.wikipedia.org/wiki/Software_testing)
- [http://martinfowler.com/tags/testing.html](https://martinfowler.com/tags/testing.html)
- [http://martinfowler.com/bliki/TestCoverage.html](https://martinfowler.com/bliki/TestCoverage.html)
- [http://martinfowler.com/bliki/TestDrivenDevelopment.html](https://martinfowler.com/bliki/TestDrivenDevelopment.html)
- [https://curl.se/docs/manpage.html](https://curl.se/docs/manpage.html)
- [https://linux.die.net/man/1/pgrep](https://linux.die.net/man/1/pgrep)
- [https://github.com/igorshubovych/markdownlint-cli](https://github.com/igorshubovych/markdownlint-cli)
- [https://go.dev/dl/](https://go.dev/dl/)
- [https://github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)
- [https://go.dev/pkg/io/#WriteString](https://pkg.go.dev/io#WriteString)
- [https://golangdocs.com/blank-identifier-in-golang](https://golangdocs.com/blank-identifier-in-golang)
- [https://blog.golang.org/error-handling-and-go](https://go.dev/blog/error-handling-and-go)
- [https://www.gnu.org/software/make/manual/html_node/Automatic-Prerequisites.html](https://www.gnu.org/software/make/manual/html_node/Automatic-Prerequisites.html)
- [https://docs.python.org/3/installing/index.html](https://docs.python.org/3/installing/index.html)
- [https://nodejs.org/en/download/](https://nodejs.org/en/download)
- [https://docs.npmjs.com/downloading-and-installing-node-js-and-npm/](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
- [https://github.com/hs-hq/W3C-Validator](https://github.com/hs-hq/W3C-Validator)
- [https://www.npmjs.com/package/markdown-link-check](https://www.npmjs.com/package/markdown-link-check)
- [https://github.com/igorshubovych/markdownlint-cli](https://github.com/igorshubovych/markdownlint-cli)
- [https://golangdocs.com/blank-identifier-in-golang](https://golangdocs.com/blank-identifier-in-golang)
- [https://go.dev/pkg/io/#WriteString](https://pkg.go.dev/io#WriteString)
- [https://github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint)

--------------------
## Tasks


### [0. Build an Application using Make](https://github.com/MathieuMorel62/holbertonschool-validation/tree/main/module2_task0)

Let’s start with a simple HTTP web server written in the [Golang](https://go.dev) language.
  
This HTTP webserver is expected to listen to incoming HTTP requests on `localhost:9999`:
  
- If a request hits the path `/` (e.g. `http://localhost:9999/` or `http://localhost:9999`) then you expect a `404 Not Found` (the homepage is not implemented in this task)
- If a request hits the path `/health` (e.g. `http://localhost:9999/health`) then you expect the server to answer `ALIVE` if it is up and running
  
With [Golang](https://go.dev) installed in your environment, initialize a new project:

```bash
## Create an empty directory and go into
mkdir awesome-api
cd awesome-api

## Initialize the Golang project with a custom identifier
go mod init github.com/<your github handle>/awesome-api
```
  
Then, create a file named `main.go` with the following content which is the Golang source code of the application:

```go
package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "os"

  "github.com/gorilla/mux"
)

func main() {
  httpAddr := "0.0.0.0:9999"
  if port := os.Getenv("PORT"); port != "" {
    httpAddr = "0.0.0.0:" + port
  }
  fmt.Println("HTTP Server listening on", httpAddr)

  // Start an HTTP server using the custom router
  log.Fatal(http.ListenAndServe(httpAddr, setupRouter()))
}

func setupRouter() *mux.Router {
  // Create a new empty HTTP Router
  r := mux.NewRouter()

  // When an HTTP GET request is received on the path /health, delegates to the function "HealthCheckHandler()"
  r.HandleFunc("/health", HealthCheckHandler).Methods("GET")

  return r
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
  // Print a line in the logs
  fmt.Println("HIT: healthcheck")

  // Write the string "ALIVE" into the response's body
  io.WriteString(w, "ALIVE")

  // End of the function: return HTTP 200 by default
}
```

#### Project Life-cycle

The life-cycle of this project is the following:
  
- “build”: compile the source code of the application to a binary named `awesome-api` (the name `awesome-api` comes from the command `go mod init github.com/<your github handle>/awesome-api`) with the command `go build`. The first build may takes some times.

- “run”: Run the application in background by executing the binary `awesome-api`, and write logs into a file named `awesome-api.log` with the command `./awesome-api >./awesome-api.log 2>&1 &`.

- “stop”: Stop the application with the command `kill XXXXX` where `XXXXX` is the Process ID of the application. For instance: `kill "$(pgrep awesome-api)"`.

- “clean”: Stop the application. Delete the binary `awesome-api` and the log file `awesome-api.log`

- “test”: You want to test it to ensure that it behaves as expected. With the application started, you may want to use the command line `curl` (or your web browser, or the command `wget` or any other HTTP client):

```bash
➜ curl http://localhost:9999
404 Page Not Found
➜ curl http://localhost:9999/health
ALIVE
```
   
Please note that you should have a line `HIT: healtcheck` in the application’s log (file `./awesome-api.log`) for each time you make a request to `http://localhost:9999/health`.

#### Requirements

You are expected to write a `Makefile` to automate the life-cycle of this application:

- A `Makefile` should be present and valid
- The binary `awesome-api` must NOT exist at the beginning, in the source code
- The goals `build`, `run`, `stop`, `clean` `test` should be implemented and mapped to the life-cycle stages of the same name:

```makefile
➜ ls -1 ./awesome-api ./awesome-api.log
ls: ./awesome-api: No such file or directory
ls: ./awesome-api.log: No such file or directory

➜ make build
➜ ls -1 ./awesome-api ./awesome-api.log
ls: ./awesome-api.log: No such file or directory
./awesome-api

➜ make test
# ...
curl: (7) Failed to connect to localhost port 9999: Connection refused
# ...
curl: (7) Failed to connect to localhost port 9999: Connection refused

➜ make run
➜ ls -1 ./awesome-api ./awesome-api.log
./awesome-api
./awesome-api.log
➜ cat ./awesome-api.log
HTTP Server listening on localhost:9999

➜ make test
# ...
404 page not found
# ...
ALIVE
➜ cat ./awesome-api.log
HTTP Server listening on localhost:9999
HIT: healtcheck

➜ make stop
➜ make test
# ...
curl: (7) Failed to connect to localhost port 9999: Connection refused
# ...
curl: (7) Failed to connect to localhost port 9999: Connection refused

➜ make clean
➜ ls -1 ./awesome-api ./awesome.log
ls: ./awesome-api: No such file or directory
ls: ./awesome-api.log: No such file or directory
```

- The goal `make help` must be implemented and print a list of all the goals with a sentence each:

```makefile
➜ make help
help: ...
build: ...
clean: ...
run: ....
```

- ⚠️ Do not forget the documentation: `README.md` file, Makefile commands, `make help` target ?
 
--------------------

## AUTHOR

- Mathieu Morel
