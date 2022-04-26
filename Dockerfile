# syntax=docker/dockerfile:1.2

FROM golang as go-build-stage

WORKDIR /go/chisel-web-proxy

RUN apt update -y && apt install -y protobuf-compiler
ADD go.mod go.sum /go/chisel-web-proxy/

ADD proto proto
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && protoc --go_out=. --go-grpc_out=. --proto_path=proto proto/*.proto

RUN go mod download

COPY cmd /go/chisel-web-proxy/cmd
COPY internal /go/chisel-web-proxy/internal

ENV GOPATH=/go PATH=$GOPATH/bin:$PATH GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=on

RUN go build -o /app/ ./...

#####################################################################

FROM node:lts as vue-build-stage

WORKDIR /app

ENV EXTEND_ESLINT=true

RUN apt update -y && apt install -y protobuf-compiler && ls -lha

COPY pkg/frontend/package*.json ./
RUN npm install \
    && npm install -g @vue/cli \
    && npm install -g protoc-gen-grpc-web

COPY proto proto/
  
RUN mkdir -p src/proto \
    && protoc -I proto proto/*.proto --js_out=import_style=commonjs:./src/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/proto

COPY pkg/frontend/ . 
RUN ls -lha proto
RUN vue build

#####################################################################

FROM nginx

WORKDIR /

COPY --from=go-build-stage /app/ /app/backend/
COPY --from=vue-build-stage /app/dist /app/frontend/
COPY scripts/ /
COPY scripts/resources/nginx-default.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080 8081 9000

RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]