# <p align=center>Holbertonschool Validation</p>

<img src="https://www.freecodecamp.org/news/content/images/size/w2000/2021/12/2220.jpg" width="100%">

# Reproducing Production Environment Locally

This project contains a shell script to reproduce a pseudo production environment locally. It simulates the steps taken by the production team to build the website in an Ubuntu 18.04 environment.

## Getting Started

To use this script, you need to have Docker installed on your machine.

### Prerequisites

- Docker

### Usage

1. Clone this repository to your local machine.
2. Open a terminal and navigate to the project's root directory (where the `setup.sh` script is located).
3. Run the following command to start an Ubuntu 18.04 Docker container:

```bash
docker run --rm --tty --interactive --volume=$(pwd):/app --workdir=/app ubuntu:18.04 /bin/bash
```

4. Inside the Docker container, execute the `./setup.sh` script. This will update the package list, install Hugo and Make, and then run the `make build` command, which should produce an error as described in the exercise.

5. When you're done, you can exit the Docker container by typing `exit`.

### Test

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

-------------------

## Author

- Mathieu Morel
