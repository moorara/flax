# BUILD STAGE
FROM golang:1.11-alpine as builder
RUN apk add --no-cache git
WORKDIR /workspace
COPY . .
RUN ./scripts/build.sh --main main.go --binary flax

# FINAL STAGE
FROM alpine:3.8
EXPOSE 8080 8443
HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD wget -q -O - http://localhost:8080/liveness || exit 1
RUN apk add --no-cache ca-certificates
COPY --from=builder /workspace/flax /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/flax
USER nobody
ENTRYPOINT [ "flax" ]
