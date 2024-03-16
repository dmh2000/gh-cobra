#!/bin/sh
go build .
echo REST API : gh-cobra api octocat
./gh-cobra api octocat
echo "==============================================="
echo GRAPHQL manual : gh-cobra graphql octocat
echo "==============================================="
./gh-cobra graphql octocat
echo "==============================================="
echo GRAPHQL shurcool : gh-cobra shurcool octocat
echo "==============================================="
./gh-cobra shurcool octocat
echo "==============================================="
echo GRAPHQL gogithub : gh-cobra gogithub octocat
echo "==============================================="
./gh-cobra shurcool octocat
echo "==============================================="
echo LS, PWD, DU : gh-cobra explain ls pwd du
echo "==============================================="
./gh-cobra explain ls pwd du
