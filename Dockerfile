FROM golang:1.10

WORKDIR /go/src/discord-bot
COPY . .

RUN go get -d -v ./... \
    && go install -v ./...

EXPOSE 80

CMD ["discord-bot"]
