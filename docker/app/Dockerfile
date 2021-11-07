FROM golang:1.16-alpine

WORKDIR /go/src
COPY ./src .

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/swaggo/gin-swagger
RUN go get -u github.com/swaggo/files

CMD ["air", "-c", ".air.toml"]
