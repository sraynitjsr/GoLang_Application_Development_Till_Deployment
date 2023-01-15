FROM golang:1.19

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /my-app

EXPOSE 8080

CMD [ "/my-app" ]
