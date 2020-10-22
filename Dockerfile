FROM golang:alpine AS builder

WORKDIR $GOPATH/src/app
COPY ./src/go-hpa ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch

COPY --from=builder /app .

EXPOSE 8000

ENTRYPOINT ["./app"]