FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /powServer

EXPOSE 9797

CMD [ "/powServer" ]