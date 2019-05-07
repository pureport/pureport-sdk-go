LINTIGNOREDEPS='vendor/.+\.go'
LINTIGNORESWAGGER='pureport/client/.+\.go'
LINTIGNOREENCRYPTION='pureport/encryption/.+\.go'

# Public SDK Core packages
SDK_CORE_PKGS=./pureport/...
SDK_COMPA_PKGS=${SDK_CORE_PKGS}

# SDK additional packages
SDK_EXAMPLES_PKGS=./example/...
SDK_TESTING_PKGS=./testing/...
SDK_MODELS_PKGS=./models/...
SDK_CLIENT_PKGS=./pureport/...
SDK_ALL_PKGS=${SDK_COMPA_PKGS} ${SDK_TESTING_PKGS} ${SDK_EXAMPLES_PKGS} ${SDK_MODELS_PKGS} ${SDK_CLIENT_PKGS}

all: verify build

# --------------------------------------------------
#  Build
# --------------------------------------------------
build: unit
	GO_EXTLINK_ENABLED=0 CGO_ENABLED=0 go build -x --ldflags '-extldflags "-static"' -o pp-cli .

# --------------------------------------------------
#  Code Generation
# --------------------------------------------------
generate: cleanup-models gen-swagger

gen-swagger:
	@echo "Generating Client for swagger definition"
	go generate ./models/swagger
	rm docs/client/*
	mv pureport/client/README.md docs/client/
	mv pureport/client/docs/* docs/client/
	rmdir pureport/client/docs
	find . -iname "*.md" -exec sed -i '' 's/\[default to null\]//g' {} +

gen-strings:
	@echo "Generating strings files for enums"
	go generate ${SDK_CLIENT_PKGS}

cleanup-models:
	@echo "Cleaning up stale model versions"
	git clean -xdf ./pureport/swagger

# --------------------------------------------------
#  Unit/CI Testing
# --------------------------------------------------
unit: gen-strings verify
	@echo "go test SDK packages"
	GO_EXTLINK_ENABLED=0 CGO_ENABLED=0 go test -v --ldflags '-extldflags "-static"' ${SDK_ALL_PKGS}

unit-with-race-cover: verify
	@echo "go test SDK and vendor packages"
	go test -count=1 -tags ${UNIT_TEST_TAGS} -race -cpu=1,2,4 ${SDK_ALL_PKGS}

ci-test: generate unit-with-race-cover ci-test-generate-validate

ci-test-generate-validate:
	@echo "CI test validate no generated code changes"
	git add . -A
	gitstatus=`git diff --cached --ignore-space-change`; \
	echo "$$gitstatus"; \
	if [ "$$gitstatus" != "" ] && [ "$$gitstatus" != "skipping validation" ]; then echo "$$gitstatus"; exit 1; fi

# --------------------------------------------------
#  Sandbox Testing
# --------------------------------------------------
sandbox-tests: sandbox-test-go112

sandbox-build-go112:
	docker build -f ./pureporttesting/sandbox/Dockerfile.test.go1.12 -t "pureport-sdk-go-1.12" .
sandbox-go112: sandbox-build-go112
	docker run -i -t pureport-sdk-go-1.12 bash
sandbox-test-go112: sandbox-build-go112
	docker run -t pureport-sdk-go-1.12

# --------------------------------------------------
#  Linting/Verify
# --------------------------------------------------
verify: lint vet

lint:
	@echo "go lint SDK and vendor packages"
	@lint=`golint ./...`; \
	dolint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDEPS} -e ${LINTIGNORESWAGGER} -e ${LINTIGNOREENCRYPTION}`; \
	echo "$$dolint"; \
	if [ "$$dolint" != "" ]; then exit 1; fi

vet:
	go vet --all ${SDK_ALL_PKGS}

# --------------------------------------------------
#  Dependencies
# --------------------------------------------------
get-deps: get-deps-verify

get-deps-tests:
	@echo "go get SDK testing dependencies"
	go get github.com/stretchr/testify

get-deps-x-tests:
	@echo "go get SDK testing golang.org/x dependencies"
	go get golang.org/x/net/http2

get-deps-verify:
	@echo "go get SDK verification utilities"
	go get golang.org/x/lint/golint

# --------------------------------------------------
#  Benchmarks
# --------------------------------------------------
bench:
	@echo "go bench SDK packages"
	go test -count=1 -run NONE -bench . -benchmen -tags 'bench' ${SDK_ALL_PKGS}
