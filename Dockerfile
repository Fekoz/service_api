FROM golang:1.19 AS builder

COPY ./src /src
WORKDIR /src

RUN echo -e "machine bitbucket.org\nlogin ___\npassword ___" > /root/.netrc

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates \
		dnsutils \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /data/conf

WORKDIR /app

VOLUME /data/conf

CMD ["./service_api", "-conf", "/data/conf"]