FROM golang:1.19-alpine

WORKDIR /src
COPY . .

ENV GO111MODULE=on

RUN go build -o /bin/action ./cmd/linkcheck

ENTRYPOINT ["/bin/action"]