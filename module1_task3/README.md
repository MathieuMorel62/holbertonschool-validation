# <p align=center>Holbertonschool Validation</p>

<img src="https://camo.githubusercontent.com/65f00c4f146b2c20e3c6ae1abc7501a4937e1fb54239fcc41860da0c59d27e56/68747470733a2f2f6d69726f2e6d656469756d2e636f6d2f6d61782f333030302f312a6d717630334b726c47354c4b325855317556344c4a672e676966" width="100%">

## Introduction to DevOps: Automate Everything to Focus on What Really Matters.

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

## Lifecycle

The lifecycle of this project is organized into several key steps:

1. **Create a new post**: Use the command `make post POST_NAME="post_name" POST_TITLE="Post Title"` to create a new post file in the `content/posts/` directory with a specified name and title.

2. **Edit the content**: Open the created post file and edit its content by adding text, images, and links as needed.

3. **Build the site**: Use the command `make build` to build the website with Hugo and generate the files in the `dist/` directory. This step transforms Markdown files into HTML pages and assembles them into a complete website structure.

4. **Clean**: If you need to remove the `dist/` directory and all the files it contains, use the `make clean` command. This step is useful if you want to rebuild the site from scratch or remove temporary files.

5. **Deployment**: Once you have successfully built the site, you can deploy the generated files in the `dist/` directory to a web server or hosting service of your choice.

6. **Help**: To display the list of available commands and their descriptions, you can use the `make help` command in the terminal. This command extracts the comments associated with each target in the Makefile and displays them to give you a quick overview of the project's features.
