FROM cgr.dev/chainguard/wolfi-base:latest AS builder

RUN apk add --no-cache go~1.26

WORKDIR /build
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /workflow-test .

FROM cgr.dev/chainguard/wolfi-base:latest
COPY --from=builder /workflow-test /usr/local/bin/workflow-test
USER nonroot:nonroot
ENTRYPOINT ["workflow-test"]
