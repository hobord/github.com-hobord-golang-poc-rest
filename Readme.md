GO development in VSCode DevContainer
=====================================

This is a POC project for testing the VSCode devcontainer environment.

I using my pre-build golang Docker image: [hobord/golang-vscode](https://github.com/hobord/golang-vscode/blob/master/Dockerfile)

It is supports go modules, languageserver, debugger...

The .devcontainer contains the docker-compose.yaml, which define the golang workspace container and a mysql server container and a phpmyadmin container.

If You don't want to use multiple containers, You can switch to single container mode in the devcontainer.json config.


The go project is a single simple example of rest server for foo entity with ddd architecture.
You can use this project as boilerplate...