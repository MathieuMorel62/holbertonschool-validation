build: ## Builds the website with Hugo and generates the files
	hugo -d dist/

clean: ## Deletes all files on dist directory
	rm -rf dist/*

post: ## Creates a new publication file in the content/posts/ with a specified name and title
	hugo new posts/$(POST_NAME).md
	sed -i 's/title: .*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md

help: ## Displays the list of targets and their use
	@grep -E -i [a-z]+:.# Makefile | tr -d '##'