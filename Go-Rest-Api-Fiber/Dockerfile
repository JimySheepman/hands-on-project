# Build Stage
FROM golang:alpine as build

WORKDIR /src/api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/api

# Final Stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/api /app/

EXPOSE 8080

CMD ./api