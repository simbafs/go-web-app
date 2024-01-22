FROM debian:bullseye-slim as build

WORKDIR /build
LABEL org.opencontainers.image.source="https://github.com/simbafs/coscup-attendance"
LABEL org.opencontainers.image.authors="Simba Fs <me@simbafs.cc>"

ENV PATH="/usr/local/go/bin:/usr/local/node-v18.17.1-linux-x64/bin:$PATH"

# install dependencies
RUN apt-get update && \
    apt-get install -y make xz-utils wget ca-certificates git --no-install-recommends && \
    # config ssl
    mkdir -p /etc/ssl/certs && \
    update-ca-certificates --fresh && \
    # install go v1.20
    wget -q https://go.dev/dl/go1.20.7.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.7.linux-amd64.tar.gz && \
    # install nodejs v16 and pnpm
    wget -q https://nodejs.org/dist/v18.17.1/node-v18.17.1-linux-x64.tar.xz && \
    tar -C /usr/local -xJf node-v18.17.1-linux-x64.tar.xz && \
    npm install -g pnpm

COPY . .
RUN make dep && \
    make build 

FROM scratch 

WORKDIR /app

COPY --from=build /build/main /app/main
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 3000
CMD [ "./main" ]
