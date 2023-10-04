FROM golang:1.20 as compiler

ARG COMMIT_SHA

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommitSHA=$COMMIT_SHA" -o /bin/apod_service cmd/main.go

FROM alpine:3.16

RUN apk add --no-cache --upgrade bash tzdata && \
    apk add ca-certificates  && \
    update-ca-certificates

COPY --from=compiler /bin/apod_service /bin/apod_service

CMD ["bin/apod_service"]
