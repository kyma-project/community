# Create a common validator for Kyma

Created on 2019-01-10 by Michał Hudy (@michal-hudy).

## Status

Proposed on 2019-01-15.

## Motivation

The Kyma project is a one big mono repository with multiple separated repositories. That repositories are validated in a different ways. Current situation causes following issues:

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
 - `test` - execute tests for component
 - `lint` - execute linters for component
 - `format` - execute code formating for component
 - `build-image` - build docker image with component
 - `push-image` - push docker image to repository

 After unification, `makefile` for `ui-api-layer` component will look like that:

 ```makefile
APP_NAME = ui-api-layer
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)

.PHONY: ci-pr ci-master ci-release resolve build test format lint build-image push-image

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test format lint

resolve:
	dep ensure -v -vendor-only
build:
	go build -o $(APP_NAME)
test:
	go test ./...
format:
	go fmt
lint:
	go vet

build-image:
	docker build -t $(APP_NAME) .
push-image:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)
 ```

### Common Makefile

As in previous point unified targets were introduced it would be fine to not duplicate all that `makefile` content in other components. Fortunately it is possible, because it is possible to include `makefiles`. Thanks to that we can create one common `makefile` that will be included in components `makefiles`.

Lets name this common `makefile` `template.mk` - files with `*.mk` are threated as `makefiles` - and put it in `kyma/common` directory.

The content of `template.mk` file will look like that:

```makefile
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)

.PHONY: ci-pr ci-master ci-release resolve build test format lint build-image push-image

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test format lint

resolve:
	dep ensure -v -vendor-only
build:
	go build -o $(APP_NAME)
test:
	go test ./...
format:
	go fmt
lint:
	go vet

build-image:
	docker build -t $(APP_NAME) .
push-image:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)
```

And `makefile` for `ui-api-layer`:

```makefile
APP_NAME = ui-api-layer
REPOSITORY_PATH = $(realpath $(shell pwd)/../..)
include $(REPOSITORY_PATH)/common//template.mk
```

Thanks to that all targets from `template.mk` are available in `makefile` for `ui-api-layer`.

Unfortunately the structure of components may vary and they may be built in different way. For example `ui-api-layer` has `main.go` file on the root of the component, `binding-usage-controller` has it in `cmd/controller` directory. The template should handle such situation and it is also possible thanks to `template definition` and `eval` function.

After introducing `template definition` the `makefiles` will look like that:

`template.mk`:
```makefile
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)

.PHONY: ci-pr ci-master ci-release resolve build test format lint build-image push-image

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test format lint

define TARGETS
resolve:
	dep ensure -v -vendor-only
build:
	go build -o $(APP_NAME) $(1)
test:
	go test ./...
format:
	go fmt
lint:
	go vet
endef

build-image:
	docker build -t $(APP_NAME) .
push-image:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)
```

`ui-api-layer`:
```makefile
APP_NAME = ui-api-layer
REPOSITORY_PATH = $(realpath $(shell pwd)/../..)
include $(REPOSITORY_PATH)/common//template.mk

$(eval $(call TARGETS,main.go))
```

Thanks to `template.mk` we will have one place in repository with targets definitions and it will be much easier to maintain it.

>**NOTE**: All components needs to be build by CI system in case of changes in `template.mk`.

### Sandbox Validation

Developers often have different environment than CI system. For example different version of Golang. It is important to provide a way for verifying code against CI environment. For that lets create another target `validate-sandbox`, that will execute all verification in same `buildpack` Docker image as CI system is using.

Definition of `validate-sandbox` target in `template.mk`:
```makefile
COMPONENT_REL_PATH=$(shell echo $(shell pwd) | sed 's,$(REPOSITORY_PATH)/,,g')

validate-sandbox:
	docker run --rm -v "$(REPOSITORY_PATH):/workspace/go/src/github.com/kyma-project/kyma" \
	--workdir "/workspace/go/src/github.com/kyma-project/kyma" \
	eu.gcr.io/kyma-project/prow/test-infra/buildpack-golang:v20181119-afd3fbd \
	make -C "/workspace/go/src/github.com/kyma-project/kyma/$(COMPONENT_REL_PATH)" validate
```

## Summary

To unify our validation and `makefile` targets we should define a template that will be included in all components `makefiles`.

The definition of template `makefile` may looks like that:
```makefile
IMG = $(DOCKER_PUSH_REPOSITORY)$(DOCKER_PUSH_DIRECTORY)/$(APP_NAME)
TAG = $(DOCKER_TAG)
COMPONENT_REL_PATH=$(shell echo $(shell pwd) | sed 's,$(REPOSITORY_PATH)/,,g')

.PHONY: ci-pr ci-master ci-release resolve build test format lint build-image push-image

ci-pr: validate build-image push-image
ci-master: ci-pr
ci-release: ci-master

validate: resolve build test format lint

validate-sandbox:
	docker run --rm -v "$(REPOSITORY_PATH):/workspace/go/src/github.com/kyma-project/kyma" \
	--workdir "/workspace/go/src/github.com/kyma-project/kyma" \
	eu.gcr.io/kyma-project/prow/test-infra/buildpack-golang:v20181119-afd3fbd \
	make -C "/workspace/go/src/github.com/kyma-project/kyma/$(COMPONENT_REL_PATH)" validate

define TARGETS
resolve:
	dep ensure -v -vendor-only
build:
	go build -o $(APP_NAME) $(1)
test:
	go test ./...
format:
	go fmt
lint:
	go vet
endef

build-image:
	docker build -t $(APP_NAME) .
push-image:
	docker tag $(APP_NAME) $(IMG):$(TAG)
	docker push $(IMG):$(TAG)
```

And a component `makefile` will be simple as:
```makefile
APP_NAME = ui-api-layer
REPOSITORY_PATH = $(realpath $(shell pwd)/../..)
include $(REPOSITORY_PATH)/common/template.mk

$(eval $(call TARGETS,main.go))
```

Thanks to that we will be able to have one place per repository with targets definitions and it will be easier to introduce new validations.