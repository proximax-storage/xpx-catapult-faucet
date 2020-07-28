FROM gcr.io/distroless/base-debian10

WORKDIR /faucet
# copy required files
COPY cmd/server.run .
COPY frontend/dist dist/

# server.run as entry point
ENTRYPOINT [ "/faucet/server.run" ]

#set the config by default
CMD ["-configFile", "/faucet/resources/rest.json", "-dist", "/faucet/dist/"]
