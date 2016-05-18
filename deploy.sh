#!/usr/bin/env bash
trap 'exit 2' ERR

echo "deploy .."

# let's do something simple first. later see: https://sebest.github.io/post/using-travis-ci-to-build-docker-images/
docker login -e $DOCKER_EMAIL -u $DOCKER_USER -p $DOCKER_PASS
docker tag gameserver filipre/gameserver:latest
docker tag gameserver filipre/gameserver:travis-$TRAVIS_BUILD_NUMBER
docker push filipre/gameserver

echo "OK"
