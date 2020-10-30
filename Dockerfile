FROM ubuntu:latest

RUN mkdir /app

ADD inventory /app

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

WORKDIR /app

EXPOSE 8081

ENTRYPOINT ["./inventory", "serve"]