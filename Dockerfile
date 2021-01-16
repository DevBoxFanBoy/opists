FROM golang:1.10 AS build
WORKDIR /go/src
COPY pkg ./pkg
COPY cmd/main.go .
COPY go.mod .
COPY ui ./ui
COPY config.yml .
COPY favicon.ico .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o opists .

FROM scratch AS runtime
COPY --from=build /go/src/opists ./
COPY --from=build /go/src/config.yml ./
COPY --from=build /go/src/favicon.ico ./
COPY --from=build /go/src/ui ./ui
EXPOSE 8080/tcp
ENTRYPOINT ["./opists"]
