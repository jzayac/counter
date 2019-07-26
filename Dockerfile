FROM golang:1.12.7-alpine3.10 as builder

COPY ./ /go/src/counter
WORKDIR /go/src/counter
RUN dep ensure && \
    CGO_ENABLED=0 go build main.go

FROM scratch
COPY --from=builder /go/src/counter/main .
EXPOSE 8080
CMD ["./main"]
