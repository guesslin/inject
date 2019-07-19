DOCKERFILE = $(if ${TARGET},${TARGET},"Dockerfile")
REPO       = inject
TAG        = latest


.PHONY: help
help:
	@echo "build: build $(REPO) via Dockerfile"
	@echo "help:  print out help text"

.PHONY: build
build:
	@docker build . -f $(DOCKERFILE) -t $(REPO):$(TAG)
