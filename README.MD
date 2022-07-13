# ASCII ART WEB - DOCKERIZE

## Description
Dockerized web application

## Files included
* Dockerfile `Dockerfile`

## Prerequisites
In order to test and use the files provided you need to have [Docker](http://www.docker.com) installed and set-up.

## Usage
To load the Docker image provided(NB! - The file is to big so it doesnt work):
`docker load < ./images/ascii-web-images.tar`

To build a Docker image using Dockerfile provided:
`docker build  -t <tag> .`

To see the list of images:
`docker images`

To run a docker container:
`docker run --name ascii-container -d -p 8080:8080 ascii-web-image`
`Docker run -d --name=<name> -p 8080:8080 <image>`

To view docker working containers:
`docker ps`

To stop container 
`docker stop ascii-container`
`docker stop <id or name of container>`

To view docker working and finished containers:
`docker ps -a`

To load the Docker image provided:
`docker load < dockerimage-ascii-art-web.tar`


To test the app running in the container go to `http://localhost:8080/` to start using the app

## Implementation details
Only Go standard libraries were used in making of the app.
Docker was used to create the Docker objects.