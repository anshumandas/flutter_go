FROM golang:1.16.3-alpine3.12

WORKDIR /app

RUN apk add git gcc libc-dev && \
    go get -u -v github.com/cosmtrek/air

COPY . .

CMD ["air", "-d"]