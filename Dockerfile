FROM golang:1.12.6-alpine3.9

WORKDIR /faucet
# copy required files
COPY cmd/server.run .
COPY ./dist dist/

# server.run as entry point
ENTRYPOINT [ "/faucet/server.run" ]

#set the config by default
CMD ["-configFile", "/faucet/resources/rest.json", "-dist", "/faucet/dist/"]
