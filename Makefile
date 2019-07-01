DOCKERFILE = $(if ${TARGET},${TARGET},"Dockerfile")
REPO       = inject
TAG        = latest


.PHONY: build
build:
	@docker build . -f $(DOCKERFILE) -t $(REPO):$(TAG)
