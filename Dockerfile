FROM golang:1.19.3 as app_builder
WORKDIR /usr/src/app
COPY . .
RUN apt-get update; \
apt-get install -y wget unzip
RUN wget -P /tmp https://github.com/protocolbuffers/protobuf/releases/download/v21.9/protoc-21.9-linux-x86_64.zip
RUN unzip /tmp/protoc-21.9-linux-x86_64.zip -d ./third_party
#
RUN go mod download && go mod verify
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
#
RUN third_party/bin/protoc \
-Iapi/proto \
-Ithird_party/include \
--proto_path=api/proto \
--go_out=protobuf/transport/certificate \
--go_opt=paths=source_relative \
--go-grpc_out=protobuf/transport/certificate \
--go-grpc_opt=paths=source_relative \
api/proto/*.proto
#
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/certificate ./cmd/rpc
#############################################################################
FROM surnet/alpine-wkhtmltopdf:3.16.2-0.12.6-small
WORKDIR /app
#RUN apk add --no-cache bash
COPY --from=app_builder /usr/src/app/bin/certificate /app/certificate
#
VOLUME ["/app/certificatesstorage"]
ENV TEMPLATES_DIR="/app/certificatesstorage/templates"
ENV CERTIFICATES_DIR="/app/certificatesstorage/certificates"
#
#EXPOSE 1234/tcp
#
ENTRYPOINT ["./certificate"]
