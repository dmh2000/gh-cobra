#!/bin/sh
go build .
echo REST API
./gh-cobra api octocat
echo GRAPHQL manual
./gh-cobra graphql octocat
echo GRAPHQL client
./gh-cobra client octocat
echo LS, PWD, DU
./gh-cobra explain ls pwd du
