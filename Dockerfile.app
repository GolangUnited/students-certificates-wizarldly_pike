FROM golang:1.19.3-alpine as builder
WORKDIR /app/src
COPY go.mod go.sum .
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/certificate-server /app/src/cmd/rpc
###
FROM surnet/alpine-wkhtmltopdf:3.16.2-0.12.6-small
WORKDIR /app/bin
COPY --from=builder /app/bin/certificate-server /app/bin/certificate-server
RUN mkdir -p /storage
ENV STORAGE=local
ENTRYPOINT ["./certificate-server"]
