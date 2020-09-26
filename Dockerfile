FROM golang:1.14-alpine as dev

WORKDIR /go/src/overpass
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .
RUN go install github.com/dustinpianalto/overpass/...

CMD [ "go", "run", "cmd/overpass/main.go"]

from alpine

WORKDIR /bin

COPY --from=dev /go/bin/overpass ./overpass

CMD [ "overpass" ]
