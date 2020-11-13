FROM golang:1.10 AS build
WORKDIR /go/src
COPY src ./src
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o opists .

FROM scratch AS runtime
COPY --from=build /go/src/opists ./
EXPOSE 8080/tcp
ENTRYPOINT ["./opists"]
