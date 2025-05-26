SVR_VERSION?=$(or $(VERSION), dev)
GOVERSION?=$(shell go version)
SHORT_COMMIT_SHA=$(shell git rev-parse --verify HEAD | cut -c1-8)

## build section
REPO_PATH=github.com/oxidnova/novadm/backend
LD_FLAGS="-w -X $(REPO_PATH)/internal.Version=$(SVR_VERSION) -X '$(REPO_PATH)/internal.GOVersion=$(GOVERSION)' -X '$(REPO_PATH)/internal.Commit=$(SHORT_COMMIT_SHA)'"

GH_CI_USER:=$(GH_CI_USER)
GH_CI_TOKEN:=$(GH_CI_TOKEN)
GL_CI_USER:=$(GL_CI_USER)
GL_CI_TOKEN:=$(GL_CI_TOKEN)


.PHONY: build
build: settoken
	@GO111MODULE=on go build -ldflags $(LD_FLAGS) -o bin/novadm $(REPO_PATH)

settoken:
ifneq ('$(GH_CI_TOKEN)','')
	@git config --global url."https://$(GH_CI_USER):$(GH_CI_TOKEN)@github.com/tespkg".insteadOf "https://github.com/tespkg"
endif
ifneq ('$(GL_CI_TOKEN)','')
	@git config --global url."https://$(GL_CI_USER):$(GL_CI_TOKEN)@gitlab.com/".insteadOf "https://gitlab.com/"
endif

.PHONY: test
test: cover.out
	@go tool cover -func cover.out

.INTERMEDIATE: cover.out
cover.out:
	@go test -v -race -coverprofile $@ ./...


.PHONY: clean
clean:
	@rm -fr bin
