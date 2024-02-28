FROM golang:alpine AS gomaster
WORKDIR /go/src/release
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as gorelease
WORKDIR /app
RUN apk add tzdata
RUN mkdir -p ./frontend/public
COPY --from=gomaster /go/src/release/app .
COPY --from=gomaster /go/src/release/env-sample /app/.env
COPY --from=gomaster /go/src/release/frontend/public /app/rontend/public

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 1112
CMD ["./app"]