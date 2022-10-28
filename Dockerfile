FROM golang:1.19-alpine AS go-builder
RUN apk add build-base
WORKDIR /app/src
COPY . .
# RUN go mod init xpx-chain-faucet
RUN go mod tidy
WORKDIR /app/src/cmd
RUN CGO_LDFLAGS='-static' go build -o server.run


FROM node:14-alpine AS nodejs-builder
WORKDIR /app/src
COPY frontend .
RUN npm install
RUN npm run build


FROM alpine:3.16
RUN apk update
WORKDIR /faucet
# copy required files
COPY --from=go-builder /app/src/cmd/server.run .
COPY --from=nodejs-builder /app/src/dist dist/

# server.run as entry point
ENTRYPOINT [ "/faucet/server.run" ]

#set the config by default
CMD ["-configFile", "/faucet/resources/rest.json", "-dist", "/faucet/dist/"]
