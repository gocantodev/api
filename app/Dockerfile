FROM golang:latest

LABEL mainteiner="Gus <otnacog@gmail.com>"

WORKDIR /app

COPY ../go.mod .
COPY ../go.sum .

RUN go mod download

COPY .. .

ENV PORT 8080

RUN cd app && go build

CMD ["./app/app"]
