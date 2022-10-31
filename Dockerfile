FROM golang:alpine3.16

ENV APP_DIR $GOPATH/src/github.com/nurana88/shop-template

RUN apk add --no-cache bash make git tar curl

WORKDIR $APP_DIR
COPY . $APP_DIR

RUN go mod download

EXPOSE 3002
ENTRYPOINT ["/bin/server"]