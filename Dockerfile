FROM golang:1.14.4 as builder

WORKDIR /go/src/github.com/aki-yogiri/weather-store

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV SOARCH=amd64

RUN go build -o app server.go

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/aki-yogiri/weather-store/app /app
ENV SERVER_PORT="8080"
EXPOSE 8080
ENTRYPOINT ["/app"]
