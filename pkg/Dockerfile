FROM golang:alpine3.15 as builder
RUN mkdir /src
ADD . /src/
WORKDIR /src
RUN go build -o cat-cozy
FROM alpine
COPY --from=builder /src/static /app/static
COPY --from=builder /src/cat-cozy /app/cat-cozy
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/cat-cozy"]