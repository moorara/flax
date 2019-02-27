# BUILD STAGE
FROM golang:1.12-alpine as builder
RUN apk add --no-cache git
WORKDIR /workspace
COPY . .
ENV CGO_ENABLED=0
RUN wget -qO- https://raw.githubusercontent.com/moorara/cherry/master/scripts/install.sh | sh
RUN cherry build -cross-compile=false -binary-file=flax

# FINAL STAGE
FROM alpine:3.9
EXPOSE 8080 9999
HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD wget -qO- http://localhost:8080/liveness || exit 1
RUN apk add --no-cache ca-certificates
COPY --from=builder /workspace/flax /usr/local/bin/
RUN chown -R nobody:nogroup /usr/local/bin/flax
USER nobody
ENTRYPOINT [ "flax" ]
