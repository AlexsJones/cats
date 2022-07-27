FROM golang:alpine3.15 as builder
RUN mkdir /src
ADD . /src/
WORKDIR /src
RUN go build -o cats
FROM alpine
COPY --from=builder /src/static /app/static
COPY --from=builder /src/cats /app/cats
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/cats"]