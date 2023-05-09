# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ./ ./

RUN cd keycut && go build -o /keycut

EXPOSE 8080

CMD [ "/keycut" ]