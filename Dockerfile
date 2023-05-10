FROM golang:1.20.4-alpine3.17

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata git

WORKDIR /build

ADD go.mod .

ADD go.sum .

ARG TG_BOT_HOST
ARG TG_BOT_TOKEN
ARG DATABASE_URL
RUN touch .env && echo "TG_BOT_HOST=$TG_BOT_HOST" >> .env && echo "TG_BOT_TOKEN=$TG_BOT_TOKEN" >> .env && echo "DATABASE_URL=$DATABASE_URL" >> .env

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o main cmd/main.go

CMD ["./main"]