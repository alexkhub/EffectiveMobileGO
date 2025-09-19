FROM golang:latest as builder

WORKDIR /app
COPY  go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /example cmd/example.go

FROM alpine:latest
COPY --from=builder example /bin/example

ENTRYPOINT ["/bin/example"]
EXPOSE 8080

CMD ["./example"]