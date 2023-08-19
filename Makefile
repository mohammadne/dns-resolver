ldflags="-X main.BuildDate=$$(TZ=Asia/Tehran date '+%FT%T') -linkmode external"
output="main"

ifdef CommitRefName
    ldflags := "-X main.CommitRefName=$(CommitRefName) $(ldflags)"
endif

ifdef CommitSHA
    ldflags := "-X main.CommitSHA=$(CommitSHA) $(ldflags)"
endif

ifdef Output
    output := $(Output)
endif

build:
	go build --ldflags "$(shell echo $(ldflags))" -o $(output) ./main.go

test:
	go test -coverpkg=./... -v -coverprofile .coverage.out ./...
	go tool cover -func .coverage.out

deploy:
	helmsman -apply -f ./deployments/helmsman.yaml

undeploy:
	helmsman -destroy -f ./deployments/helmsman.yaml
