# Giai đoạn 1: build. Image này có đủ bộ công cụ Go, nặng vài trăm MB.
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags "-X main.version=${VERSION}" -o /out/app ./cmd

# Giai đoạn 2: chạy. Chỉ chép file thực thi sang một image nhỏ.
FROM alpine:3.19
RUN adduser -D -u 10001 app
USER app
COPY --from=build /out/app /app
EXPOSE 3000
ENTRYPOINT ["/app"]
