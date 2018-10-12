#!/bin/bash

set -e

export GOOS=$(go env GOOS)

testDocker() {
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
make ci-test
fi

if [[ $GOOS == "linux" ]]; then
testDocker
fi

if [[ $GOOS == "windows" ]]; then
make ci-test
fi
