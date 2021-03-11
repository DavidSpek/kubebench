FROM golang:1.16.1 AS build

WORKDIR /go/src/app
COPY . .

RUN env CGO_ENABLED=0 go build -a -o /kubebenchjob-controller

FROM scratch

COPY --from=build /kubebenchjob-controller /kubebenchjob-controller
ENTRYPOINT ./kubebenchjob-controller