# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /trading-bot-project

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /trading-bot-project

COPY --from=build /trading-bot-project ./trading-bot-project
COPY ./.env ./.env

ENTRYPOINT ["./trading-bot-project", "-production"]
