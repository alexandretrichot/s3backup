FROM golang:1.19-alpine AS build
ARG VERSION=developement
WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . .

RUN go build -ldflags="-X 'main.Version=${VERSION}'" -o /s3backup

FROM alpine
WORKDIR /

RUN apk update && apk add mongodb-tools

COPY --from=build /s3backup /s3backup

ENTRYPOINT ["/s3backup"]
