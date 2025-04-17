FROM golang:1.20 AS build
WORKDIR /go/src
COPY . .
ENV CGO_ENABLED=0
RUN go mod download
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o todoapi .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/todoapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./todoapi", "serve"]
