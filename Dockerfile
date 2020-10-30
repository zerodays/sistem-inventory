FROM golang:latest AS build

RUN mkdir /app
COPY ./ /app

WORKDIR /app

RUN make build

FROM ubuntu:latest

RUN mkdir /app

COPY --from=build ./app/inventory /app

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

WORKDIR /app

EXPOSE 8081

ENTRYPOINT ["./inventory", "serve"]