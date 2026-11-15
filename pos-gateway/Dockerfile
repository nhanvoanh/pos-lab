# Giai đoạn 1: build. Chạy trên kiến trúc của máy build (BUILDPLATFORM) và
# biên dịch chéo sang kiến trúc đích (TARGETARCH), nên một lần build ra được
# image cho cả máy Intel/AMD (amd64) lẫn Apple Silicon/Graviton (arm64).
FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
ARG VERSION=dev
ARG TARGETOS=linux
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags "-X main.version=${VERSION}" -o /out/app ./cmd

# Giai đoạn 2: chạy. Chỉ chép file thực thi sang một image nhỏ,
# chạy bằng user thường (10001) chứ không phải root.
FROM alpine:3.19
USER 10001
COPY --from=build /out/app /app
EXPOSE 3000
ENTRYPOINT ["/app"]
