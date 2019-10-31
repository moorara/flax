# BUILD STAGE
FROM golang:1.13-alpine as builder
RUN apk add --no-cache git
WORKDIR /flax
COPY . .
ENV CGO_ENABLED=0
RUN wget -qO - https://git.io/JeCX6 | sh
RUN cherry build -cross-compile=false

# FINAL STAGE
FROM alpine:3.10
EXPOSE 8080 9999
HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD wget -qO - http://localhost:8080/liveness || exit 1
RUN apk add --no-cache ca-certificates
COPY --from=builder /flax/bin/flax /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/flax
USER nobody
ENTRYPOINT [ "flax" ]
