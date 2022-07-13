# syntax=docker/dockerfile:1

#Baseimage
FROM golang:1.16-alpine

WORKDIR /app

# Get dependencies
COPY go.mod .
RUN go mod download

# Get files
COPY . .

# Build
RUN go build -o /ascii-art-web

# Metadata
LABEL meta-data.app-name="ascii-art-web"
LABEL meta-data.author.autor1="Sten (sten9911)" meta-data.author.autor2="Paavel (paavel_makarenko)"

# Declare port
EXPOSE 8080

# Run
CMD [ "/ascii-art-web" ]