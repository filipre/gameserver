# Game server

[![Join the chat at https://gitter.im/filipre/gameserver](https://badges.gitter.im/filipre/gameserver.svg)](https://gitter.im/filipre/gameserver?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Takes requests from smartphone clients and takes action or returns information

## Required for access

- Oauth token
- HTTPS
- required version of api

## Docker

```
$ docker build -t my-golang-app .
$ docker run -it --rm --name my-running-app my-golang-app
```

## Travis

let's trigger the first build with docker
