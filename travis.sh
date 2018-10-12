#!/bin/bash

set -ex

export GOOS=$(go env GOOS)

testDocker() {
echo "using docker"
docker build \
  --build-arg CODECOV_TOKEN=$CODECOV_TOKEN \
  --build-arg CI=$CI \
  --build-arg TRAVIS=$TRAVIS \
  --build-arg TRAVIS_BRANCH=$TRAVIS_BRANCH \
  --build-arg TRAVIS_COMMIT=$TRAVIS_COMMIT \
  --build-arg TRAVIS_JOB_ID=$TRAVIS_JOB_ID \
  --build-arg TRAVIS_JOB_NUMBER=$TRAVIS_JOB_NUMBER \
  --build-arg TRAVIS_OS_NAME=$TRAVIS_OS_NAME \
  --build-arg TRAVIS_PULL_REQUEST=$TRAVIS_PULL_REQUEST \
  --build-arg TRAVIS_PULL_REQUEST_SHA=$TRAVIS_PULL_REQUEST_SHA \
  --build-arg TRAVIS_REPO_SLUG=$TRAVIS_REPO_SLUG \
  --build-arg TRAVIS_TAG=$TRAVIS_TAG \
  .
}

echo $GOOS

if [[ $GOOS == "darwin" ]]; then
echo "testing for darwin"
# travis doesn't support docker for mac
go test -tags "sqlite integration_test" -race ./...
fi

if [[ $GOOS == "linux" ]]; then
echo "testing for linux"
testDocker
fi

if [[ $GOOS == "windows" ]]; then
echo "testing for windows"
go test -tags "sqlite integration_test" -race ./...
fi
