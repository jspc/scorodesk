BUILDNUMBER = 1

default: deps build

deps:
	go get -v

build:
	go build -ldflags "-X 'main.buildNumber=$(BUILDNUMBER)' \
			   -X 'main.buildMachine=$(shell hostname)' \
			   -X 'main.builtBy=$(shell whoami)' \
			   -X 'main.builtWhen=$(shell date +'%s')' \
			   -X 'main.compiler=$(shell go version)' \
			   -X 'main.sha=$(sha git rev-parse HEAD)'" \
			   -o 'scorodesk'

upload:
	s3cmd sync scorodesk s3://beamly-pipeline-artifacts/scorodesk/${GO_PIPELINE_LABEL}/
	s3cmd sync scorodesk s3://beamly-pipeline-artifacts/scorodesk/latest/
