FROM --platform=$BUILDPLATFORM golang:1.17.8-alpine3.15 as builder
RUN mkdir /build
WORKDIR /build

ARG TARGETOS
ARG TARGETARCH

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -o pi-temp .


# final image
FROM alpine:3.15
COPY --from=builder /build/pi-temp .

ENTRYPOINT [ "./pi-temp" ]