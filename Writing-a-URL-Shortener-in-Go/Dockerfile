FROM golang:1.11

WORKDIR /go/src/gitlab.com/cfilby/go-short
COPY . .

RUN go install
EXPOSE 8080
CMD ["go-short"]