FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.21-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG VERSION="dev"

RUN apk update && apk add -U --no-cache ca-certificates

WORKDIR /app/
ADD go.mod go.sum ./
RUN go mod download
ADD main.go .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -o civo-versions main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch
WORKDIR /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/civo-versions /app/civo-versions
ENTRYPOINT ["/app/civo-versions"]
