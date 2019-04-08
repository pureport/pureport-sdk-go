LINTIGNOREDEPS='vendor/.+\.go'
LINTIGNORESWAGGER='pureport/swagger/.+\.go'

# Public SDK Core packages
SDK_CORE_PKGS=./pureport/...
SDK_COMPA_PKGS=${SDK_CORE_PKGS}

# SDK additional packages
SDK_EXAMPLES_PKGS=./example/...
SDK_TESTING_PKGS=./testing/...
SDK_MODELS_PKGS=./models/...
SDK_ALL_PKGS=${SDK_COMPA_PKGS} ${SDK_TESTING_PKGS} ${SDK_EXAMPLES_PKGS} ${SDK_MODELS_PKGS}

all: generate unit

# --------------------------------------------------
#  Code Generation
# --------------------------------------------------
generate: cleanup-models gen-swagger

gen-swagger:
	@echo "Generating Client for swagger definition"
	go generate ./models/swagger

cleanup-models:
	@echo "Cleaning up stale model versions"
	git clean -xdf ./pureport/swagger

# --------------------------------------------------
#  Unit/CI Testing
# --------------------------------------------------
unit: verify

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
#  Linting/Verify
# --------------------------------------------------
verify: lint vet

lint:
	@echo "go lint SDK and vendor packages"
	@lint=`golint ./...`; \
	dolint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDEPS} -e ${LINTIGNORESWAGGER}`; \
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
	go test -count=1 -run NONE -bench . -benchmen -tags 'bench '${SDK_ALL_PKGS}
