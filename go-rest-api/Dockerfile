# TODO: Don't copy to .env file research to fix method

FROM golang:1.17 as build-env

WORKDIR /go/src/app

COPY .env ./

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=build-env /go/bin/app /
CMD ["/app"]