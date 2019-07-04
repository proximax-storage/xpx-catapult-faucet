FROM golang:1.12.6-alpine3.9

WORKDIR /faucet
# copy required files
COPY cmd/server.run .
COPY ./resources/ resources/
COPY ./dist dist/

# execute go binary
CMD ["server.run", "-configFile", "resources/rest.json", "-dist", "dist/"]
