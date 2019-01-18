# Create a common validator for Kyma

Created on 2019-01-10 by Michał Hudy (@michal-hudy).

## Status

Proposed on 2019-01-15.

## Motivation

The Kyma project is a one big mono repository with multiple separate projects. That projects are validated in a different ways. Current situation causes following issues:

- Script that validates repositories is duplicated across components.
- Multiple versions of scripts that validates components.
- It is not possible to maintain it.
- Contributors don't know what they should execute to validate component.
- Different formatting and coding standards in Kyma components.
- CI validates code in a different environment.

## Solution

Github `kyma-project` organization already contains `makefiles` in every repository and CI pipeline is based on it. In such case, we should unify `makefiles` targets across components/repositories and drop support for any other solution like before-commit, custom validators etc. Thanks to that, whenever contributor will work on `console`, `kyma component` or some other `tool` in our organization, she/he will know how to execute build, test, or format code. It is crucial for development experience.

I will base my proposal on `ui-api-layer` component and targets may not work as it is only an example.

### Makefile Targets

Makefile targets should be simple and should allow contributors to execute all required validations as one command. It should be unified across all `kyma-project` repositories regardless language or type.

Currently, `makefile` in components looks like that:

```makefile
APP_NAME = ui-api-layer
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)

.PHONY: ci-pr ci-master ci-release resolve build-and-test build-image push-image

ci-pr: resolve build-and-test build-image push-image
ci-master: resolve build-and-test build-image push-image
ci-release: resolve build-and-test build-image push-image

resolve:
	dep ensure -v -vendor-only
build-and-test:
	./before-commit.sh ci
build-image:
	docker build -t $(APP_NAME) .
push-image:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)
```

Only the `ci-*` targets are unified in `kyma-project` organization, because they are required by CI system. Other targets may have different names, because they are not unified and not required by any system. For example `ui-api-layer` has `build-and-test` target, `binding-usage-controller` has `build` - both targets executes `before-commit.sh` script. In such case contributor doesn't know how to work with components in one repository.

As a solution for that I would like to unify targets:
 - `ci-pr`, `ci-master`, `ci-release` - targets required by CI system
 - `validate` - execute all validation
 - `resolve` - download dependencies
 - `build` - execute build of component
 - `clean` - remove artifacts
 - `test` - execute tests for component
 - `lint` - execute linters for component
 - `format` - execute code formating for component
 - `validate-format` - validates if files are formated correctly
 - `docker-build` - build docker image with component
 - `docker-push` - push docker image to repository

 After unification, `makefile` for `ui-api-layer` component will look like that:

 ```makefile
APP_NAME = ui-api-layer
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)

.PHONY: ci-pr ci-master ci-release validate resolve build test format lint build-image push-image

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test format lint

resolve:
	dep ensure -v -vendor-only
build:
	go build -o bin/$(APP_NAME)
clean:
	rm bin/$(APP_NAME)
test:
	go test ./...
format:
	go fmt ./...
validate-format:
	$(eval CHANGED:=$(shell go fmt ./...))
	@test -z "$(CHANGED)" \
	|| (echo "Not formatted files: $(CHANGED)" && exit 1)
lint:
	go vet ./...

docker-build:
	docker build -t $(APP_NAME) .
docker-push:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)
 ```

### Common Makefile

As in previous point unified targets were introduced it would be fine to not duplicate all that `makefile` content in other components. Fortunately it is possible to include `makefiles`. Thanks to that we can create one common `makefile` that will be included in components `makefiles`.

Lets name this common `makefile` `template.go.mk` - files with `*.mk` are treated as `makefiles` - and put it in `kyma/scripts` directory. In `kyma/scripts` we can also store other common scripts.

The content of `template.go.mk` file will look like that:

```makefile
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)

.PHONY: ci-pr ci-master ci-release validate resolve build test format lint build-image push-image

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test format lint

resolve:
	dep ensure -v -vendor-only
build:
	go build -o bin/$(APP_NAME)
clean:
	rm bin/$(APP_NAME)
test:
	go test ./...
format:
	go fmt ./...
validate-format:
	$(eval CHANGED:=$(shell go fmt ./...))
	@test -z "$(CHANGED)" \
	|| (echo "Not formatted files: $(CHANGED)" && exit 1)
lint:
	go vet ./...

docker-build:
	docker build -t $(APP_NAME) .
docker-push:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)
```

And `makefile` for `ui-api-layer`:

```makefile
APP_NAME = ui-api-layer
REPOSITORY_PATH = $(realpath $(shell pwd)/../..)

include $(REPOSITORY_PATH)/scripts/template.go.mk
```

Thanks to that all targets from `template.go.mk` are available in `makefile` for `ui-api-layer`.

Unfortunately the structure of components may vary and they may be built in different way. For example `ui-api-layer` has `main.go` file on the root of the component, `binding-usage-controller` has it in `cmd/controller` directory. The template should handle such situation and it is also possible thanks to `template definition` and `eval` function.

After introducing `template definition` the `makefiles` will look like that:

`template.go.mk`:
```makefile
DOCKER_REPOSITORY = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY) # provided by CI system
COMPONENT_REL_PATH=$(shell echo $(shell pwd) | sed 's,$(REPOSITORY_PATH)/,,g')

.PHONY: ci-pr ci-master ci-release resolve validate build clean test format validate-format lint docker-build docker-push

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test validate-format lint clean
resolve:
	dep ensure -v -vendor-only
test:
	go test ./...
format:
	go fmt ./... # may be replaced by goimports
validate-format:
	$(eval CHANGED:=$(shell go fmt ./...))
	@test -z "$(CHANGED)" \
	|| (echo "Not formatted files: $(CHANGED)" && exit 1)
lint:
	go vet ./...

define TARGETS
build:
	go build -o bin/$(APP_NAME) $(1)
clean):
	rm bin/$(APP_NAME)
docker-build:
	docker build -t $(APP_NAME) . --file $(2)
docker-push:
	docker tag $(APP_NAME) $(DOCKER_REPOSITORY)/$(APP_NAME):$(DOCKER_TAG)
	docker push $(DOCKER_REPOSITORY)/$(APP_NAME):$(DOCKER_TAG)
endef
```

`ui-api-layer`:
```makefile
APP_NAME = ui-api-layer
REPOSITORY_PATH = $(realpath $(shell pwd)/../..)

include $(REPOSITORY_PATH)/scripts/template.go.mk

$(eval $(call TARGETS,main.go,Dockerfile))
```

Thanks to `template.go.mk` we will have one place in repository with targets definitions and it will be much easier to maintain it.

>**NOTE**: All components needs to be build by CI system in case of changes in `template.go.mk`.

### Sandbox Validation

Developers often have different environment than CI system. For example different version of Golang. It is important to provide a way for verifying code against CI environment. For that lets create targets with `-sandbox` suffix, that will executes targets in same `buildpack` Docker image as CI system is using.

To achieve that we need to update `template.go.mk` and `Makefile` in component

`template.go.mk`:
```makefile
DOCKER_REPOSITORY = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY) # provided by CI system
COMPONENT_REL_PATH=$(shell echo $(shell pwd) | sed 's,$(REPOSITORY_PATH)/,,g')

.PHONY: ci-pr ci-master ci-release resolve validate build clean test format validate-format lint docker-build docker-push

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test validate-format lint clean
resolve:
	dep ensure -v -vendor-only
test:
	go test ./...
format:
	go fmt ./... # may be replaced by goimports
validate-format:
	$(eval CHANGED:=$(shell go fmt ./...))
	@test -z "$(CHANGED)" \
	|| (echo "Not formatted files: $(CHANGED)" && exit 1)
lint:
	go vet ./...

define TARGETS
build:
	go build -o bin/$(APP_NAME) $(1)
clean):
	rm bin/$(APP_NAME)
docker-build:
	docker build -t $(APP_NAME) . --file $(2)
docker-push:
	docker tag $(APP_NAME) $(DOCKER_REPOSITORY)/$(APP_NAME):$(DOCKER_TAG)
	docker push $(DOCKER_REPOSITORY)/$(APP_NAME):$(DOCKER_TAG)
endef

define SANDBOX
.PHONY: $(1)-sandbox
$(1)-sandbox:
	docker run --rm -v "$(REPOSITORY_PATH):/workspace/go/src/github.com/kyma-project/kyma" \
	--workdir "/workspace/go/src/github.com/kyma-project/kyma/$(COMPONENT_REL_PATH)" \
	eu.gcr.io/kyma-project/prow/test-infra/buildpack-golang:$(BUILDPACK_VERSION) \
	make $(1)
endef

$(eval $(call SANDBOX,validate))
$(eval $(call SANDBOX,validate-format))
$(eval $(call SANDBOX,resolve))
$(eval $(call SANDBOX,test))
$(eval $(call SANDBOX,format))
$(eval $(call SANDBOX,lint))
```

To component `Makefile` we need to add `buildpack` version:
```makefile
APP_NAME = ui-api-layer
BUILDPACK_VERSION = v20181119-afd3fbd
REPOSITORY_PATH = $(realpath $(shell pwd)/../..)

include $(REPOSITORY_PATH)/scripts/template.go.mk

$(eval $(call TARGETS,main.go,Dockerfile))
```

### Support for multiple artifacts from one component

Some components like `event-bus` generates multiple artifacts during build. Handling such situation is also possible, but it will requires unification of components structure. All `main.go` files will need to be located in `cmd/{name}/main.go` and names of `Dockerfiles` will need to be changed to `{name}.Dockerfile` or be in separated directories.

We should investigate what will be easier to achieve - one artifact per component or unification structure of components.

`template.go.mk` that supports multiple artifacts:
```makefile
DOCKER_REPOSITORY = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY) # provided by CI system
COMPONENT_REL_PATH=$(shell echo $(shell pwd) | sed 's,$(REPOSITORY_PATH)/,,g')

.PHONY: ci-pr ci-master ci-release resolve validate build clean test format validate-format lint docker-build docker-push

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test validate-format lint clean
resolve:
	dep ensure -v -vendor-only
test:
	go test ./...
format:
	go fmt ./... # may be replaced by goimports
validate-format:
	$(eval CHANGED:=$(shell go fmt ./...))
	@test -z "$(CHANGED)" \
	|| (echo "Not formatted files: $(CHANGED)" && exit 1)
lint:
	go vet ./...
build: $(foreach appName,$(APP_NAMES),build-$(appName))
clean: $(foreach appName,$(APP_NAMES),clean-$(appName))
docker-build: $(foreach appName,$(APP_NAMES),docker-build-$(appName))
docker-push: $(foreach appName,$(APP_NAMES),docker-push-$(appName))

define TARGETS
.PHONY: build-$(1) clean-$(1) docker-build-$(1) docker-push-$(1)
build-$(1):
	go build -o bin/$(1) cmd/$(1)
clean-$(1):
	rm bin/$(1)
docker-build-$(1):
	docker build -t $(1) . --file $(1).Dockerfile
docker-push-$(1):
	docker tag $(1) $(DOCKER_REPOSITORY)/$(1):$(DOCKER_TAG)
	docker push $(DOCKER_REPOSITORY)/$(1):$(DOCKER_TAG)
endef

define SANDBOX
.PHONY: $(1)-sandbox
$(1)-sandbox:
	docker run --rm -v "$(REPOSITORY_PATH):/workspace/go/src/github.com/kyma-project/kyma" \
	--workdir "/workspace/go/src/github.com/kyma-project/kyma/$(COMPONENT_REL_PATH)" \
	eu.gcr.io/kyma-project/prow/test-infra/buildpack-golang:$(BUILDPACK_VERSION) \
	make $(1)
endef

$(eval $(call SANDBOX,validate))
$(eval $(call SANDBOX,validate-format))
$(eval $(call SANDBOX,resolve))
$(eval $(call SANDBOX,test))
$(eval $(call SANDBOX,format))
$(eval $(call SANDBOX,lint))
```

And `Makefile` in component:
```makefile
APP_NAMES = name1 name2 name3 name4
BUILDPACK_VERSION = v20181119-afd3fbd
REPOSITORY_PATH = $(realpath $(shell pwd)/../..)

include $(REPOSITORY_PATH)/scripts/template.go.mk

$(foreach appName,$(APP_NAMES),$(eval $(call TARGETS,$(appName))))
```

## Summary

To unify our validation and `Makefile` targets we should define a template that will be included in all components `Makefiles`. Thanks to that we will be able to have one place per repository with targets definitions and it will be easier to introduce new validations.

Such proposal also introduces the possibility to validate source code against CI environment, what can speed up creating Pull Requests. It is possible to support multiple artifacts per components, but we should consider if we still want it.

Examples:
 - [One artifact per component](#sandbox-validation)
 - [Multiple artifacts per component](#support-for-multiple-artifacts-from-one-component)

The disadvantages of this solution:
 - Template is defined per repository, not per organization.
 - `buildpack` version is defined per component and is duplicated with CI definition.
 - It is not visible what targets are available in component `makefile`, because they are generated during execution.
