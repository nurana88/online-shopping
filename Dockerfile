# FROM is instruction for base image. It creates operating system. Basically it is important to create infrastructure of container
# alpine means golang has very basic programs. alpine versions are lighter
FROM golang:alpine3.16

ENV APP_DIR $GOPATH/src/github.com/nurana88/shop-template

# apk is the Alpine Package Keeper - the distribution's package manager.
# It is the primary method for installing additional software, and is available in the apk-tools package
RUN apk add --no-cache bash make git tar curl

WORKDIR $APP_DIR
COPY . $APP_DIR

RUN go mod download

# Expose is instruction for aws to map incomming traffic in 3002. It's not for developers but for aws cloud
EXPOSE 3002
ENTRYPOINT ["/bin/server"]