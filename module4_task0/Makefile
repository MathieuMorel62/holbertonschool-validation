.PHONY: help build run stop lint test unit-tests integration-tests check validate clean post

HUGO=hugo

lint: ## Run markdown linting on the DEPLOY.md and README.md files
	markdownlint -f DEPLOY.md README.md || true

build: ## Build the Hugo website and create a placeholder file for the Go application
	$(HUGO)
	touch awesome-api

run: ## Execute the placeholder Go application and redirect output to awesome-api.log
	./awesome-api >./awesome-api.log 2>&1 &

stop: ## Terminate the running placeholder Go application
	@pkill awesome-api

clean: ## Stop the application and remove the placeholder Go application, log file, coverage files, and Hugo website build
	@rm -rf awesome-api awesome-api.log coverage-units.out coverage-integrations.out dist/ *~

test: ## Run both unit and integration tests for the Go application
	make unit-tests
	make integration-tests
	make validate

help: ## Displays the list of targets and their use
	@grep -E -i ^[a-z] Makefile | tr -d '##' | sort

unit-tests: ## Run unit tests, integration tests, and HTML validation
	@hugo
	touch coverage-units.out

integration-tests: ## Run integration tests for the Hugo website and generate an integration test coverage file
	@hugo
	touch coverage-integrations.out

check: ## Run markdown linting on blog posts and check for dead links in all markdown files
	markdownlint ./content/posts/*.md
	find content -name \*.md -print0 | xargs -0 -n1 markdown-link-check

validate: ## Validate the generated HTML file using Holberton's W3C Validator
	python3 ./w3c_validator.py ./dist/index.html || true

post: ## Create a new blog post and set its title
	hugo new posts/$(POST_NAME).md
	sed -i 's/title: .*/title: "$(POST_TITLE)"/' content/posts/$(POST_NAME).md

package: ## Build the Hugo website and create a ZIP file containing the website build and placeholder Go application
	@hugo
	zip -FSr awesome-website.zip . -i dist/ awesome-api
