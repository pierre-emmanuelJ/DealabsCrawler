FROM golang:1.11.5-alpine
MAINTAINER Pierre-Emmanuel Jacquier <pierre-emmanuel.jacquier@epitech.eu>

WORKDIR /go/src/github.com/pierre-emmanuelJ/DealabsCrawler
COPY . .
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dealabscrawler .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0  /go/src/github.com/pierre-emmanuelJ/DealabsCrawler/dealabscrawler .
CMD ["./dealabscrawler", "--sender-mail", "example@gmail.com", "--sender-mail-password", "passwordExample"]
