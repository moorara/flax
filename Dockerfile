# BUILD STAGE
FROM golang:1.15-alpine as builder
RUN apk add --no-cache git
WORKDIR /repo
COPY . .
RUN wget -qO - https://git.io/JeCX6 | sh
RUN cherry build -cross-compile=false

# FINAL STAGE
FROM alpine:3.13
EXPOSE 8080 9999
RUN apk add --no-cache ca-certificates
COPY --from=builder /repo/bin/flax /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/flax
USER nobody
ENTRYPOINT [ "flax" ]
